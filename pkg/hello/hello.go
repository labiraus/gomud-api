package hello

import (
	"context"
	"log"
	"net"

	pb "github.com/labiraus/gomud-api/api/exported"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":8080"
)

// server is used to implement helloworld.HelloServer.
type server struct {
	pb.UnimplementedHelloServer
}

// SayHello implements helloworld.HelloServer
func (s *server) SayHello(ctx context.Context, request *emptypb.Empty) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "from heck off"}, nil
}

func Start(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		close(done)
	}
	s := grpc.NewServer()
	go func() {
		defer close(done)
		pb.RegisterHelloServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	go func() {
		<-ctx.Done()
		s.Stop()
	}()
	return done
}
