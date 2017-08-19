package commands

import (
	"flag"
	"os"
	"os/signal"
	"time"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/nicholasjackson/bench/server"
	"github.com/nicholasjackson/bench/server/proto"
)

// Run is a cli command which allows running of benchmarks
type Run struct {
	flagSet *flag.FlagSet

	pluginLocation string
	threads        int
	duration       time.Duration
	rampUp         time.Duration
	timeout        time.Duration

	client *plugin.Client
}

// Help returns the command help
func (r *Run) Help() string {
	r.flagSet.Usage()
	return ""
}

// Run runs the command
func (r *Run) Run(args []string) int {
	r.flagSet.Parse(args)

	if r.pluginLocation != "" {
		req := proto.RunRequest{
			PluginLocation: r.pluginLocation,
			Threads:        int64(r.threads),
			Duration:       int64(r.duration),
			Ramp:           int64(r.rampUp),
			Timeout:        int64(r.timeout),
		}

		client := server.NewGRPCClient()

		go func(cl *server.GRPCClient) {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)

			switch <-c {
			case os.Interrupt:
				cl.Stop()
			}
		}(client)

		client.Run(req)
	}

	return 0
}

// Synopsis returns information about the command
func (r *Run) Synopsis() string {
	return "run the benchmarks"
}

// NewRun creates a new Run command
func NewRun() *Run {
	r := &Run{
		flagSet: flag.NewFlagSet("run", flag.ContinueOnError),
	}

	r.flagSet.StringVar(&r.pluginLocation, "plugin", "", "specify the location of the bench plugin")
	r.flagSet.IntVar(&r.threads, "thread", 1, "the number of concurrent threads when running a benchmark")
	r.flagSet.DurationVar(&r.duration, "duration", 10*time.Second, "the duration of the test e.g. 5s (5 seconds)")
	r.flagSet.DurationVar(&r.rampUp, "ramp", 10*time.Second, "time taken to schedule maximum threads e.g. 5s (5 seconds)")
	r.flagSet.DurationVar(&r.timeout, "timeout", 5*time.Second, "timeout value for a test e.g. 5s (5 seconds)")

	return r
}
