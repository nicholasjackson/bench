package server

import (
	"log"

	"github.com/nicholasjackson/bench/server/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client proto.BenchServerClient
}

func (g *GRPCClient) Run(r proto.RunRequest) {
	_, err := g.client.Run(context.Background(), &r)
	if err != nil {
		log.Println(err)
	}
}

func (g *GRPCClient) Stop() {
	_, err := g.client.Stop(context.Background(), &proto.ServerEmpty{})
	if err != nil {
		log.Println(err)
	}
}

func (g *GRPCClient) StartPlugin(pluginLocation string) {
	_, err := g.client.StartPlugin(
		context.Background(),
		&proto.StartPluginRequest{PluginLocation: pluginLocation},
	)

	if err != nil {
		log.Println(err)
	}
}

func (g *GRPCClient) Execute() error {
	_, err := g.client.Execute(context.Background(), &proto.ExecuteRequest{})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (g *GRPCClient) Close() {
	g.conn.Close()
}

func NewGRPCClient(uri string) *GRPCClient {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewBenchServerClient(conn)

	return &GRPCClient{client: c, conn: conn}
}
