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
	Uuid      string
	Type      string
	Role      string
}

type TokenClaimsResult struct {
	Uuid string
	Type string
	Role string
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
		return fmt.Errorf("unable to connect to '%v', error: %w", s.serverHost, err)
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
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not verify token, error during connection: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.VerifyToken(ctx, &VerifyTokenRequest{Token: token})
	if err != nil {
		return nil, fmt.Errorf("could not get verify token; token: '%v'; error: %w", token, err)
	}

	result := ToVerificationResult(reply)

	return &result, nil
}

func (s *AuthGRPCService) GetTokenClaims(token string) (*TokenClaimsResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not get token claims; error during connection: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetTokenClaims(ctx, &GetTokenClaimsRequest{Token: token})
	if err != nil {
		return nil, fmt.Errorf("could not get token claims; token: '%v'; error: %w", token, err)
	}

	result := ToTokenClaimsResult(reply)

	return &result, nil
}

func ToVerificationResult(reply *VerifyTokenReply) VerificationResult {
	return VerificationResult{
		IsValid:   reply.GetIsValid(),
		IsExpired: reply.GetIsExpired(),
		Uuid:      reply.GetUuid(),
		Type:      reply.GetType(),
		Role:      reply.GetRole(),
	}
}

func ToTokenClaimsResult(reply *GetTokenClaimsReply) TokenClaimsResult {
	return TokenClaimsResult{
		Uuid: reply.GetUuid(),
		Type: reply.GetType(),
		Role: reply.GetRole(),
	}
}
