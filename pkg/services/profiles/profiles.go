package profiles

import (
	context "context"
	"fmt"
	"io"
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

func (s *ProfilesGRPCService) GetUsers(offset int32, limit int32, shard int32) (*GetUsersReply, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetUsers: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetUsers(ctx, &GetUsersRequest{Offset: offset, Limit: limit, Shard: shard})
	if err != nil {
		return nil, fmt.Errorf("could not GetUsers: %v", err)
	}
	return reply, nil
}

func (s *ProfilesGRPCService) GetUsersByUuids(userUuids []string) (*GetUsersReply, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetUsers: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetUsersByUuids(ctx, &GetUsersByUuidsRequest{Uuids: userUuids})
	if err != nil {
		return nil, fmt.Errorf("could not GetUsers: %v", err)
	}
	return reply, nil
}

func (s *ProfilesGRPCService) GetUsersStream(userUuids []string) (<-chan (GetUserResult), error) {
	var result chan (GetUserResult) = make(chan GetUserResult)
	var resultErr error
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetUsersStream: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	stream, err := s.client.GetUsersStream(ctx)
	if err != nil {
		return result, fmt.Errorf("could not GetUsersStream: %v", err)
	}
	defer stream.CloseSend()

	if err != nil {
		return result, fmt.Errorf("s.client.GetUsersStream failed: %v", err)
	}
	go func() {
		defer close(result)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				resultErr = fmt.Errorf("s.client.GetUsersStream: stream.Recv failed: %v", err)
				return
			}
			result <- ToGetUserResult(in)
		}
	}()
	for _, userUuid := range userUuids {
		err := stream.Send(&GetUserRequest{Uuid: userUuid})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetUsersStream: stream.Send(%v) failed: %v", userUuid, err)
			return result, resultErr
		}
	}
	return result, resultErr
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

func ToGetGetUserResultSlice(replies []*GetUserReply) []GetUserResult {
	result := []GetUserResult{}

	for _, p := range replies {
		result = append(result, ToGetUserResult(p))
	}

	return result
}
