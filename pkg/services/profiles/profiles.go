package profiles

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CredentialsValidationResult struct {
	UserId  int  `json:"userId" binding:"required"`
	IsValid bool `json:"isValid" binding:"required"`
}

type ProfilesGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       ProfilesServiceClient
	queryTimeout time.Duration
}

func CreateProfilesGRPCService(serverHost string) *ProfilesGRPCService {
	// TODO: add TLS
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return &ProfilesGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *ProfilesGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewProfilesServiceClient(conn)
	return nil
}

func (s *ProfilesGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *ProfilesGRPCService) ValidateCredentials(login string, password string) (*CredentialsValidationResult, error) {
	var result *CredentialsValidationResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.ValidateCredentials(ctx, &ValidateCredentialsRequest{Login: login, Password: password})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	result = &CredentialsValidationResult{
		IsValid: reply.GetIsValid(),
		UserId:  int(reply.GetUserId()),
	}

	return result, nil
}
