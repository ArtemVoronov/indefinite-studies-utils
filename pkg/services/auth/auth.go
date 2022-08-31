package auth

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type VerificationResult struct {
	IsValid   bool
	IsExpired bool
}

type TokenClaimsResult struct {
	Id   int
	Type string
}

type AuthGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       AuthServiceClient
	queryTimeout time.Duration
}

func CreateAuthGRPCService(serverHost string, creds *credentials.TransportCredentials) *AuthGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &AuthGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *AuthGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewAuthServiceClient(conn)
	return nil
}

func (s *AuthGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *AuthGRPCService) VerifyToken(token string) (*VerificationResult, error) {
	var result *VerificationResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.VerifyToken(ctx, &VerifyTokenRequest{Token: token})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	result = &VerificationResult{
		IsValid:   reply.GetIsValid(),
		IsExpired: reply.GetIsExpired(),
	}

	return result, nil
}

func (s *AuthGRPCService) GetTokenClaims(token string) (*TokenClaimsResult, error) {
	var result *TokenClaimsResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetTokenClaims(ctx, &GetTokenClaimsRequest{Token: token})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	result = &TokenClaimsResult{
		Id:   int(reply.GetId()),
		Type: reply.GetType(),
	}

	return result, nil
}
