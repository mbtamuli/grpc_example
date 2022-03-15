package forwarder

import (
	"mriyam.dev/grpc-example/helloworld"
	pb "mriyam.dev/grpc-example/helloworld"
)

func HandleSendToStream(stream helloworld.Greeter_SayHelloServer, reply *pb.HelloReply, done chan error) {
	if err := stream.Send(reply); err != nil {
		done <- err
	}

	done <- nil
}
