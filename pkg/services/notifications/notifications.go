package notifications

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type NotificationsGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       NotificationsServiceClient
	queryTimeout time.Duration
}

func CreateNotificationsGRPCService(serverHost string, creds *credentials.TransportCredentials) *NotificationsGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &NotificationsGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *NotificationsGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewNotificationsServiceClient(conn)
	return nil
}

func (s *NotificationsGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *NotificationsGRPCService) SendEmail(sender string, recepient string, body string) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not SendEmail: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.SendEmail(ctx, &SendEmailRequest{Sender: sender, Recepient: recepient, Body: body})
	if err != nil {
		return fmt.Errorf("could not SendEmail: %v", err)
	}

	return nil
}
