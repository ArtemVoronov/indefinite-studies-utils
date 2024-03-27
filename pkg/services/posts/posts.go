package posts

import (
	context "context"
	"fmt"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GetPostResult struct {
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

type GetCommentResult struct {
	Id              int
	AuthorUuid      string
	PostUuid        string
	LinkedCommentId string
	Text            string
	State           string
	CreateDate      time.Time
	LastUpdateDate  time.Time
}

type GetTagResult struct {
	Id   int
	Name string
}

type PostsGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       PostsServiceClient
	queryTimeout time.Duration
}

func CreatePostsGRPCService(serverHost string, creds *credentials.TransportCredentials) *PostsGRPCService {
	var opts []grpc.DialOption
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &PostsGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *PostsGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %w", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewPostsServiceClient(conn)
	return nil
}

func (s *PostsGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *PostsGRPCService) GetPost(postUuid string) (*GetPostResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetPost: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetPost(ctx, &GetPostRequest{Uuid: postUuid})
	if err != nil {
		return nil, fmt.Errorf("could not GetPost: %w", err)
	}

	result := ToGetPostResult(reply)

	return &result, nil
}

func (s *PostsGRPCService) GetComment(postUuid string, commentId int64) (*GetCommentResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetComment: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetComment(ctx, &GetCommentRequest{PostUuid: postUuid, Id: commentId})
	if err != nil {
		return nil, fmt.Errorf("could not GetComment: %w", err)
	}

	result := ToGetCommentResult(reply)

	return &result, nil
}

func (s *PostsGRPCService) GetTag(id int64) (*GetTagResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetTag: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetTag(ctx, &GetTagRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("could not GetTag: %w", err)
	}

	result := ToGetTagResult(reply)

	return &result, nil
}

func (s *PostsGRPCService) GetTags(offset int32, limit int32) (*GetTagsReply, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetTags: %w", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetTags(ctx, &GetTagsRequest{Offset: offset, Limit: limit})
	if err != nil {
		return nil, fmt.Errorf("could not GetTags: %w", err)
	}
	return reply, nil
}

func ToGetPostResult(post *GetPostReply) GetPostResult {
	return GetPostResult{
		Uuid:           post.GetUuid(),
		AuthorUuid:     post.GetAuthorUuid(),
		Text:           post.GetText(),
		PreviewText:    post.GetPreviewText(),
		Topic:          post.GetTopic(),
		State:          post.GetState(),
		CreateDate:     post.GetCreateDate().AsTime(),
		LastUpdateDate: post.GetLastUpdateDate().AsTime(),
		TagIds:         utils.ToInt(post.TagIds),
	}
}

func ToGetCommentResult(comment *GetCommentReply) GetCommentResult {
	return GetCommentResult{
		Id:              int(comment.GetId()),
		AuthorUuid:      comment.GetAuthorUuid(),
		PostUuid:        comment.GetPostUuid(),
		LinkedCommentId: comment.GetLinkedCommentId(),
		Text:            comment.GetText(),
		State:           comment.GetState(),
		CreateDate:      comment.GetCreateDate().AsTime(),
		LastUpdateDate:  comment.GetLastUpdateDate().AsTime(),
	}
}

func ToGetPostsResultSlice(posts []*GetPostReply) []GetPostResult {
	result := []GetPostResult{}

	for _, p := range posts {
		result = append(result, ToGetPostResult(p))
	}

	return result
}

func ToGetTagResult(tag *GetTagReply) GetTagResult {
	return GetTagResult{
		Id:   int(tag.Id),
		Name: tag.Name,
	}
}

func ToGetTagResultSlice(tags []*GetTagReply) []GetTagResult {
	result := []GetTagResult{}

	for _, p := range tags {
		result = append(result, ToGetTagResult(p))
	}

	return result
}
