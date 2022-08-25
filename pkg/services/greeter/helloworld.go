package greeter

import (
	context "context"
	"fmt"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type HelloService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       GreeterClient
	queryTimeout time.Duration
}

func CreateHelloService(serverHost string, creds *credentials.TransportCredentials) *HelloService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &HelloService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *HelloService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewGreeterClient(conn)
	return nil
}

func (s *HelloService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *HelloService) Hello(name string) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		return fmt.Errorf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", reply.GetMessage())

	return nil
}
