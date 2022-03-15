// Package main implements a server for Greeter service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"mriyam.dev/grpc-example/forwarder"
	"mriyam.dev/grpc-example/helloworld"
	pb "mriyam.dev/grpc-example/helloworld"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(in *pb.HelloRequest, stream helloworld.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", in.GetName())

	done := make(chan error, 1)
	reply := &pb.HelloReply{Message: "Hello " + in.GetName()}
	go forwarder.HandleSendToStream(stream, reply, done)

	return <-done
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
