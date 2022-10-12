package feed

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedPostDTO struct {
	Uuid           string
	AuthorId       int32
	Text           string
	PreviewText    string
	Topic          string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
	Tags           []string
}

type FeedCommentDTO struct {
	Id              int32
	Uuid            string
	AuthorId        int32
	PostUuid        string
	LinkedCommentId int32
	Text            string
	State           string
	CreateDate      time.Time
	LastUpdateDate  time.Time
}

type FeedUserDTO struct {
	Id             int32
	Login          string
	Email          string
	Role           string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
}

type FeedBuilderGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       FeedBuilderServiceClient
	queryTimeout time.Duration
}

func CreateFeedBuilderGRPCService(serverHost string, creds *credentials.TransportCredentials) *FeedBuilderGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &FeedBuilderGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *FeedBuilderGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewFeedBuilderServiceClient(conn)
	return nil
}

func (s *FeedBuilderGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *FeedBuilderGRPCService) CreatePost(post *FeedPostDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not CreatePost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.CreatePost(ctx, toCreatePostRequest(post))
	return err
}

func (s *FeedBuilderGRPCService) UpdatePost(post *FeedPostDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not UpdatePost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.UpdatePost(ctx, toUpdatePostRequest(post))
	return err
}

func (s *FeedBuilderGRPCService) DeletePost(postUuid string) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not DeletePost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.DeletePost(ctx, &DeletePostRequest{Uuid: postUuid})
	return err
}

func (s *FeedBuilderGRPCService) CreateComment(comment *FeedCommentDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not CreateComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.CreateComment(ctx, toCreateCommentRequest(comment))
	return err
}

func (s *FeedBuilderGRPCService) UpdateComment(comment *FeedCommentDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not UpdateComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.UpdateComment(ctx, toUpdateCommentRequest(comment))
	return err
}

func (s *FeedBuilderGRPCService) DeleteComment(postUuid string, commentUuid string) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not DeleteComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.DeleteComment(ctx, &DeleteCommentRequest{Uuid: commentUuid, PostUuid: postUuid})
	return err
}
func (s *FeedBuilderGRPCService) UpdateUser(user *FeedUserDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not UpdateUser: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.UpdateUser(ctx, toUpdateUserRequest(user))
	return err
}

func toCreatePostRequest(post *FeedPostDTO) *CreatePostRequest {
	return &CreatePostRequest{
		Uuid:           post.Uuid,
		AuthorId:       post.AuthorId,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
		Tags:           post.Tags,
	}
}

func toUpdatePostRequest(post *FeedPostDTO) *UpdatePostRequest {
	return &UpdatePostRequest{
		Uuid:           post.Uuid,
		AuthorId:       post.AuthorId,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
		Tags:           post.Tags,
	}
}

func toCreateCommentRequest(comment *FeedCommentDTO) *CreateCommentRequest {
	return &CreateCommentRequest{
		Id:              int32(comment.Id),
		Uuid:            comment.Uuid,
		AuthorId:        comment.AuthorId,
		PostUuid:        comment.PostUuid,
		LinkedCommentId: comment.LinkedCommentId,
		Text:            comment.Text,
		State:           comment.State,
		CreateDate:      timestamppb.New(comment.CreateDate),
		LastUpdateDate:  timestamppb.New(comment.LastUpdateDate),
	}
}

func toUpdateCommentRequest(comment *FeedCommentDTO) *UpdateCommentRequest {
	return &UpdateCommentRequest{
		Id:              int32(comment.Id),
		Uuid:            comment.Uuid,
		AuthorId:        comment.AuthorId,
		PostUuid:        comment.PostUuid,
		LinkedCommentId: comment.LinkedCommentId,
		Text:            comment.Text,
		State:           comment.State,
		CreateDate:      timestamppb.New(comment.CreateDate),
		LastUpdateDate:  timestamppb.New(comment.LastUpdateDate),
	}
}

func toUpdateUserRequest(user *FeedUserDTO) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:    user.Id,
		Login: user.Login,
		Email: user.Email,
		Role:  user.Role,
		State: user.State,
	}
}
