package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/nicholasjackson/bench/bench"
	"github.com/nicholasjackson/bench/bench/output"
	"github.com/nicholasjackson/bench/bench/util"
	"github.com/nicholasjackson/bench/plugin/shared"
	"github.com/nicholasjackson/bench/server/proto"
)

var benchProcess *bench.Bench

type GRPCServer struct{}

func (g *GRPCServer) Execute(context.Context, *proto.ExecuteRequest) (*proto.ServerEmpty, error) {
	log.Println("Execute")
	return nil, nil
}

func (g *GRPCServer) Run(c context.Context, r *proto.RunRequest) (*proto.ServerEmpty, error) {
	log.Println("Running Bench")

	runBench(r.PluginLocation, r.Threads, time.Duration(r.Duration), time.Duration(r.Ramp), time.Duration(r.Timeout))

	return &proto.ServerEmpty{}, nil
}

func (g *GRPCServer) Stop(c context.Context, r *proto.ServerEmpty) (*proto.ServerEmpty, error) {
	benchProcess.Stop()
	return &proto.ServerEmpty{}, nil
}

func NewGRPCServer() {
	log.Println("Starting Bench Server")
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	proto.RegisterBenchServerServer(s, &GRPCServer{})

	log.Fatal(s.Serve(lis))
}

func createPlugin(pluginLocation string) (*plugin.Client, shared.Bench) {
	logger := hclog.New(&hclog.LoggerOptions{
		Output: hclog.DefaultOutput,
		Level:  hclog.Info,
		Name:   "plugin",
	})

	c := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", pluginLocation),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           logger,
	})

	// Connect via RPC
	grpcClient, err := c.Client()
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	plug, err := grpcClient.Dispense("bench")
	if err != nil {
		fmt.Println("Error getting plugin:", err.Error())
		os.Exit(1)
	}

	return c, plug.(shared.Bench)
}

func runBench(location string, threads int64, duration time.Duration, ramp time.Duration, timeout time.Duration) {
	p, bp := createPlugin(location)
	defer p.Kill()

	benchProcess = bench.New(int(threads), duration, ramp, timeout)

	benchProcess.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	benchProcess.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	benchProcess.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	benchProcess.AddOutput(0*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)

	benchProcess.RunBenchmarks(func() error {
		return bp.Do()
	})
}
