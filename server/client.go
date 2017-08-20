package server

import (
	"log"

	"github.com/nicholasjackson/bench/server/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCClient is a gRPC client for interacting with the bench server
type GRPCClient struct {
	conn   *grpc.ClientConn
	client proto.BenchServerClient
}

// Run a new benchmark test
func (g *GRPCClient) Run(r proto.RunRequest) {
	_, err := g.client.Run(context.Background(), &r)
	if err != nil {
		log.Println(err)
	}
}

// Stop a running bench test
func (g *GRPCClient) Stop() {
	_, err := g.client.Stop(context.Background(), &proto.ServerEmpty{})
	if err != nil {
		log.Println(err)
	}
}

// StartPlugin starts a bench plugin
func (g *GRPCClient) StartPlugin(plugin []byte) {
	_, err := g.client.StartPlugin(
		context.Background(),
		&proto.StartPluginRequest{Plugin: plugin},
	)

	if err != nil {
		log.Println(err)
	}
}

// Execute a test with the bench plugin
func (g *GRPCClient) Execute() error {
	_, err := g.client.Execute(context.Background(), &proto.ExecuteRequest{})
	if err != nil {
		return err
	}

	return nil
}

// Close the client connection
func (g *GRPCClient) Close() {
	g.conn.Close()
}

// NewGRPCClient creates a new client with the given uri
func NewGRPCClient(uri string) *GRPCClient {
	conn, err := grpc.Dial(uri,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(20*1024*1024)),
	)
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewBenchServerClient(conn)

	return &GRPCClient{client: c, conn: conn}
}
