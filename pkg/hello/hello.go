package hello

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/labiraus/gomud-common/proto/gomud-api"
	user "github.com/labiraus/gomud-common/proto/gomud-user"
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
	message, err := callUser("fred")
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: message}, nil
}

func callUser(username string) (string, error) {
	conn, err := grpc.Dial("http://service-user.gomud:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &user.GreetingRequest{Name: username})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return "", nil
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
