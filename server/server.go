package server

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/hashicorp/serf/cmd/serf/command/agent"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/nicholasjackson/bench/bench"
	"github.com/nicholasjackson/bench/bench/output"
	"github.com/nicholasjackson/bench/bench/util"
	"github.com/nicholasjackson/bench/plugin/shared"
	"github.com/nicholasjackson/bench/server/proto"
	"github.com/nicholasjackson/ultraclient"
)

var benchProcess *bench.Bench
var plug *plugin.Client
var benchClient shared.Bench

type GRPCServer struct {
	serf *agent.Agent
	port int
}

func (g *GRPCServer) Execute(context.Context, *proto.ExecuteRequest) (*proto.ServerEmpty, error) {
	return &proto.ServerEmpty{}, benchClient.Do()
}

func (g *GRPCServer) StartPlugin(c context.Context, pr *proto.StartPluginRequest) (*proto.ServerEmpty, error) {
	log.Println("Start Plugin")
	plug, benchClient = createPlugin(pr.PluginLocation)

	return &proto.ServerEmpty{}, nil
}

func (g *GRPCServer) Run(c context.Context, r *proto.RunRequest) (*proto.ServerEmpty, error) {
	log.Println("Running Bench")

	// get server members
	members := make([]string, 0)
	for _, m := range g.serf.Serf().Members() {
		server := m.Addr.String() + ":" + m.Tags["benchServerPort"]
		members = append(members, server)

		log.Println("Start Plugin:", server)
		client := NewGRPCClient(server)
		defer client.Close()
		client.StartPlugin(r.PluginLocation)
	}

	runBench(r.PluginLocation, members, r.Threads, time.Duration(r.Duration), time.Duration(r.Ramp), time.Duration(r.Timeout))

	return &proto.ServerEmpty{}, nil
}

func (g *GRPCServer) Stop(c context.Context, r *proto.ServerEmpty) (*proto.ServerEmpty, error) {
	benchProcess.Stop()
	return &proto.ServerEmpty{}, nil
}

func NewGRPCServer(serf *agent.Agent, port int) {
	log.Println("Starting Bench Server on Port:", port)
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	proto.RegisterBenchServerServer(s, &GRPCServer{serf: serf})

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

func runBench(location string, servers []string, threads int64, duration time.Duration, ramp time.Duration, timeout time.Duration) {
	benchProcess = bench.New(int(threads), duration, ramp, timeout)

	benchProcess.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	benchProcess.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	benchProcess.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	benchProcess.AddOutput(0*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)

	// get a load balancer
	var endpoints []url.URL
	for _, s := range servers {
		endpoints = append(endpoints, url.URL{Host: s})
	}

	lb := ultraclient.RoundRobinStrategy{}
	bs := ultraclient.ExponentialBackoff{}

	config := ultraclient.Config{
		Timeout:                timeout,
		MaxConcurrentRequests:  500,
		ErrorPercentThreshold:  25,
		DefaultVolumeThreshold: 10,
		Retries:                1,
		Endpoints:              endpoints,
	}

	uc := ultraclient.NewClient(config, &lb, &bs)

	benchProcess.RunBenchmarks(func() error {
		return uc.Do(func(endpoint url.URL) error {
			client := NewGRPCClient(endpoint.Host)
			defer client.Close()

			return client.Execute()
		})
	})
}
