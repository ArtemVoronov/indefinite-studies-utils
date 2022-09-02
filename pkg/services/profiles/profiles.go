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
	UserId  int  `json:"userId" binding:"required"`
	IsValid bool `json:"isValid" binding:"required"`
}

type GetUserResult struct {
	Id             int
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

func (s *ProfilesGRPCService) GetUser(userId int32) (*GetUserResult, error) {
	var result *GetUserResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetUser(ctx, &GetUserRequest{Id: userId})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	result = &GetUserResult{
		Id:             int(reply.GetId()),
		Login:          reply.GetLogin(),
		Email:          reply.GetEmail(),
		Role:           reply.GetRole(),
		State:          reply.GetState(),
		CreateDate:     reply.GetCreateDate().AsTime(),
		LastUpdateDate: reply.GetLastUpdateDate().AsTime(),
	}

	return result, nil
}

func (s *ProfilesGRPCService) GetUsers(userIds []int32) ([]GetUserResult, error) {
	var result []GetUserResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetUsers(ctx, &GetUsersRequest{Ids: userIds})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	users := reply.GetUsers()

	for _, userPtr := range users {
		result = append(result, GetUserResult{
			Id:             int(userPtr.GetId()),
			Login:          userPtr.GetLogin(),
			Email:          userPtr.GetEmail(),
			Role:           userPtr.GetRole(),
			State:          userPtr.GetState(),
			CreateDate:     userPtr.GetCreateDate().AsTime(),
			LastUpdateDate: userPtr.GetLastUpdateDate().AsTime(),
		})
	}

	return result, nil
}

func (s *ProfilesGRPCService) GetUsersStream(userIds []int32) (<-chan (GetUserResult), error) {
	var result chan (GetUserResult) = make(chan GetUserResult)
	var resultErr error
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	stream, err := s.client.GetUsersStream(ctx)
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
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
			result <- GetUserResult{
				Id:             int(in.GetId()),
				Login:          in.GetLogin(),
				Email:          in.GetEmail(),
				Role:           in.GetRole(),
				State:          in.GetState(),
				CreateDate:     in.GetCreateDate().AsTime(),
				LastUpdateDate: in.GetLastUpdateDate().AsTime(),
			}
		}
	}()
	for _, userId := range userIds {
		err := stream.Send(&GetUserRequest{Id: userId})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetUsersStream: stream.Send(%v) failed: %v", userId, err)
			return result, resultErr
		}
	}
	return result, resultErr
}
