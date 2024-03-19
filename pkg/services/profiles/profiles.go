package profiles

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type CredentialsValidationResult struct {
	UserUuid string `json:"userUuid" binding:"required"`
	IsValid  bool   `json:"isValid" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type GetUserResult struct {
	Id             int
	Uuid           string
	Login          string
	Email          string
	Role           string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
}

type ProfilesGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       ProfilesServiceClient
	queryTimeout time.Duration
}

func CreateProfilesGRPCService(serverHost string, creds *credentials.TransportCredentials) *ProfilesGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

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
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not ValidateCredentials: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.ValidateCredentials(ctx, &ValidateCredentialsRequest{Login: login, Password: password})
	if err != nil {
		return nil, fmt.Errorf("could not ValidateCredentials: %v", err)
	}

	result := ToCredentialsValidationResult(reply)

	return &result, nil
}

func (s *ProfilesGRPCService) GetUser(userUuid string) (*GetUserResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetUser: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetUser(ctx, &GetUserRequest{Uuid: userUuid})
	if err != nil {
		return nil, fmt.Errorf("could not GetUser: %v", err)
	}

	result := ToGetUserResult(reply)

	return &result, nil
}

func ToCredentialsValidationResult(reply *ValidateCredentialsReply) CredentialsValidationResult {
	return CredentialsValidationResult{
		IsValid:  reply.GetIsValid(),
		UserUuid: reply.GetUserUuid(),
		Role:     reply.GetRole(),
	}
}

func ToGetUserResult(reply *GetUserReply) GetUserResult {
	return GetUserResult{
		Id:             int(reply.GetId()),
		Uuid:           reply.GetUuid(),
		Login:          reply.GetLogin(),
		Email:          reply.GetEmail(),
		Role:           reply.GetRole(),
		State:          reply.GetState(),
		CreateDate:     reply.GetCreateDate().AsTime(),
		LastUpdateDate: reply.GetLastUpdateDate().AsTime(),
	}
}
