package server

import (
	"log"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/nicholasjackson/bench/server/proto"
)

type GRPCServer struct{}

func (g *GRPCServer) Execute(context.Context, *proto.ExecuteRequest) (*proto.ServerEmpty, error) {
	log.Println("Execute")
	return nil, nil
}

func (g *GRPCServer) Run(context.Context, *proto.RunRequest) (*proto.ServerEmpty, error) {
	log.Println("Run")
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
