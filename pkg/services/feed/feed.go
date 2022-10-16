package feed

import (
	context "context"
	"fmt"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedPostDTO struct {
	Uuid           string
	AuthorUuid     string
	Text           string
	PreviewText    string
	Topic          string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
	TagIds         []int
}

type FeedCommentDTO struct {
	Id                int32
	Uuid              string
	AuthorUuid        string
	PostUuid          string
	LinkedCommentUuid string
	Text              string
	State             string
	CreateDate        time.Time
	LastUpdateDate    time.Time
}

type FeedUserDTO struct {
	Uuid           string
	Login          string
	Email          string
	Role           string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
}

type FeedTagDTO struct {
	Id   int
	Name string
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

	_, err := s.client.CreatePost(ctx, ToCreatePostRequest(post))
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

	_, err := s.client.UpdatePost(ctx, ToUpdatePostRequest(post))
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

	_, err := s.client.CreateComment(ctx, ToCreateCommentRequest(comment))
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

	_, err := s.client.UpdateComment(ctx, ToUpdateCommentRequest(comment))
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

	_, err := s.client.UpdateUser(ctx, ToUpdateUserRequest(user))
	return err
}
func (s *FeedBuilderGRPCService) CreateTag(tag *FeedTagDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not CreateTag: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.CreateTag(ctx, ToCreateTagRequest(tag))
	return err
}

func (s *FeedBuilderGRPCService) UpdateTag(tag *FeedTagDTO) error {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return fmt.Errorf("could not UpdateTag: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	_, err := s.client.UpdateTag(ctx, ToUpdateTagRequest(tag))
	return err
}

func ToCreatePostRequest(post *FeedPostDTO) *CreatePostRequest {
	return &CreatePostRequest{
		Uuid:           post.Uuid,
		AuthorUuid:     post.AuthorUuid,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
		TagIds:         utils.ToInt32(post.TagIds),
	}
}

func ToUpdatePostRequest(post *FeedPostDTO) *UpdatePostRequest {
	return &UpdatePostRequest{
		Uuid:           post.Uuid,
		AuthorUuid:     post.AuthorUuid,
		Text:           post.Text,
		PreviewText:    post.PreviewText,
		Topic:          post.Topic,
		State:          post.State,
		CreateDate:     timestamppb.New(post.CreateDate),
		LastUpdateDate: timestamppb.New(post.LastUpdateDate),
		TagIds:         utils.ToInt32(post.TagIds),
	}
}

func ToCreateCommentRequest(comment *FeedCommentDTO) *CreateCommentRequest {
	return &CreateCommentRequest{
		Id:                int32(comment.Id),
		Uuid:              comment.Uuid,
		AuthorUuid:        comment.AuthorUuid,
		PostUuid:          comment.PostUuid,
		LinkedCommentUuid: comment.LinkedCommentUuid,
		Text:              comment.Text,
		State:             comment.State,
		CreateDate:        timestamppb.New(comment.CreateDate),
		LastUpdateDate:    timestamppb.New(comment.LastUpdateDate),
	}
}

func ToUpdateCommentRequest(comment *FeedCommentDTO) *UpdateCommentRequest {
	return &UpdateCommentRequest{
		Id:                int32(comment.Id),
		Uuid:              comment.Uuid,
		AuthorUuid:        comment.AuthorUuid,
		PostUuid:          comment.PostUuid,
		LinkedCommentUuid: comment.LinkedCommentUuid,
		Text:              comment.Text,
		State:             comment.State,
		CreateDate:        timestamppb.New(comment.CreateDate),
		LastUpdateDate:    timestamppb.New(comment.LastUpdateDate),
	}
}

func ToUpdateUserRequest(user *FeedUserDTO) *UpdateUserRequest {
	return &UpdateUserRequest{
		Uuid:  user.Uuid,
		Login: user.Login,
		Email: user.Email,
		Role:  user.Role,
		State: user.State,
	}
}

func ToCreateTagRequest(tag *FeedTagDTO) *CreateTagRequest {
	return &CreateTagRequest{
		Id:   int32(tag.Id),
		Name: tag.Name,
	}
}

func ToUpdateTagRequest(tag *FeedTagDTO) *UpdateTagRequest {
	return &UpdateTagRequest{
		Id:   int32(tag.Id),
		Name: tag.Name,
	}
}

func ToTagIds(dto []FeedTagDTO) []int32 {
	result := make([]int32, 0, len(dto))
	for _, tag := range dto {
		result = append(result, int32(tag.Id))
	}
	return result
}
