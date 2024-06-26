package subscriptions

import (
	context "context"
	"fmt"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/kafka"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type SubscriptionsGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       SubscriptionsServiceClient
	queryTimeout time.Duration
}

func CreateSubscriptionsGRPCService(serverHost string, creds *credentials.TransportCredentials) *SubscriptionsGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &SubscriptionsGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *SubscriptionsGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %w", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewSubscriptionsServiceClient(conn)
	return nil
}

func (s *SubscriptionsGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *SubscriptionsGRPCService) PutEvent(eventType string, eventBody string) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not PutEvent: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.PutEvent(ctx, &PutEventRequest{EventType: eventType, EventBody: eventBody})
	if err != nil {
		return fmt.Errorf("could not PutEvent: %w", err)
	}

	return nil
}

func (s *SubscriptionsGRPCService) PutSendEmailEvent(in kafka.SendEmailEvent) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not PutSendEmailEvent: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.PutSendEmailEvent(ctx, &PutSendEmailEventRequest{
		Sender:    in.Sender,
		Recepient: in.Recepient,
		Subject:   in.Subject,
		Body:      in.Body,
	})
	if err != nil {
		return fmt.Errorf("could not PutSendEmailEvent: %w", err)
	}

	return nil
}
