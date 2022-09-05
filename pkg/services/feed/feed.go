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
	Id             int32
	AuthorId       int32
	Text           string
	PreviewText    string
	Topic          string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
}

type FeedCommentDTO struct {
	Id              int32
	AuthorId        int32
	PostId          int32
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

func (s *FeedBuilderGRPCService) DeletePost(postId int32) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not DeletePost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.DeletePost(ctx, &DeletePostRequest{Id: postId})
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

func (s *FeedBuilderGRPCService) DeleteComment(postId int32, commentId int32) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not DeleteComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.DeleteComment(ctx, &DeleteCommentRequest{Id: commentId, PostId: postId})
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

func (s *FeedBuilderGRPCService) DeleteUser(userId int32) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not DeleteUser: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.DeleteUser(ctx, &DeleteUserRequest{Id: userId})
	return err
}

func toCreatePostRequest(post *FeedPostDTO) *CreatePostRequest {
	return &CreatePostRequest{
		Id:             post.Id,
		AuthorId:       post.AuthorId,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
	}
}

func toUpdatePostRequest(post *FeedPostDTO) *UpdatePostRequest {
	return &UpdatePostRequest{
		Id:             post.Id,
		AuthorId:       post.AuthorId,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
	}
}

func toCreateCommentRequest(comment *FeedCommentDTO) *CreateCommentRequest {
	return &CreateCommentRequest{
		Id:              comment.Id,
		AuthorId:        comment.AuthorId,
		PostId:          comment.PostId,
		LinkedCommentId: comment.LinkedCommentId,
		Text:            comment.Text,
		State:           comment.State,
		CreateDate:      timestamppb.New(comment.CreateDate),
		LastUpdateDate:  timestamppb.New(comment.LastUpdateDate),
	}
}

func toUpdateCommentRequest(comment *FeedCommentDTO) *UpdateCommentRequest {
	return &UpdateCommentRequest{
		Id:              comment.Id,
		AuthorId:        comment.AuthorId,
		PostId:          comment.PostId,
		LinkedCommentId: comment.LinkedCommentId,
		Text:            comment.Text,
		State:           comment.State,
		CreateDate:      timestamppb.New(comment.CreateDate),
		LastUpdateDate:  timestamppb.New(comment.LastUpdateDate),
	}
}

func toUpdateUserRequest(user *FeedUserDTO) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:             user.Id,
		Login:          user.Login,
		Email:          user.Email,
		Role:           user.Role,
		State:          user.State,
		CreateDate:     timestamppb.New(user.CreateDate),
		LastUpdateDate: timestamppb.New(user.LastUpdateDate),
	}
}
