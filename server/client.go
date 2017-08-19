package server

import (
	"log"

	"github.com/nicholasjackson/bench/server/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	client proto.BenchServerClient
}

func (g *GRPCClient) Run(r proto.RunRequest) {
	_, err := g.client.Run(context.Background(), &r)
	if err != nil {
		log.Println(err)
	}
}

func NewGRPCClient() *GRPCClient {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewBenchServerClient(conn)

	return &GRPCClient{client: c}
}
