package posts

import (
	context "context"
	"fmt"
	"io"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/utils"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GetPostResult struct {
	Uuid           string
	AuthorId       int
	Text           string
	PreviewText    string
	Topic          string
	State          string
	CreateDate     time.Time
	LastUpdateDate time.Time
}

type GetCommentResult struct {
	Id              int
	Uuid            string
	AuthorId        int
	PostUuid        string
	LinkedCommentId *int
	Text            string
	State           string
	CreateDate      time.Time
	LastUpdateDate  time.Time
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
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
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
			return nil, fmt.Errorf("could not GetPost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetPost(ctx, &GetPostRequest{Uuid: postUuid})
	if err != nil {
		return nil, fmt.Errorf("could not GetPost: %v", err)
	}

	result := ToGetPostsResult(reply)

	return &result, nil
}

func (s *PostsGRPCService) GetPosts(offset int32, limit int32, shard int32) (*GetPostsReply, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetPosts: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetPosts(ctx, &GetPostsRequest{Offset: offset, Limit: limit, Shard: shard})
	if err != nil {
		return nil, fmt.Errorf("could not GetPosts: %v", err)
	}
	reply.GetPosts()
	return reply, nil
}

func (s *PostsGRPCService) GetPostsStream(postUuids []string) (<-chan (GetPostResult), error) {
	var result chan (GetPostResult) = make(chan GetPostResult)
	var resultErr error
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetPostsStream: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	stream, err := s.client.GetPostsStream(ctx)
	if err != nil {
		return result, fmt.Errorf("could not GetPostsStream: %v", err)
	}
	defer stream.CloseSend()

	if err != nil {
		return result, fmt.Errorf("s.client.GetPostsStream failed: %v", err)
	}
	go func() {
		defer close(result)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				resultErr = fmt.Errorf("s.client.GetPostsStream: stream.Recv failed: %v", err)
				return
			}
			result <- ToGetPostsResult(in)
		}
	}()
	for _, postUuid := range postUuids {
		err := stream.Send(&GetPostRequest{Uuid: postUuid})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetPostsStream: stream.Send(%v) failed: %v", postUuid, err)
			return result, resultErr
		}
	}
	return result, resultErr
}

func (s *PostsGRPCService) GetComment(postUuid string, commentid int32) (*GetCommentResult, error) {
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return nil, fmt.Errorf("could not GetComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetComment(ctx, &GetCommentRequest{PostUuid: postUuid, Id: commentid})
	if err != nil {
		return nil, fmt.Errorf("could not GetComment: %v", err)
	}

	result := ToGetCommentResult(reply)

	return &result, nil
}

func (s *PostsGRPCService) GetComments(postUuid string, offset int32, limit int32) ([]GetCommentResult, error) {
	var result []GetCommentResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetComments: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetComments(ctx, &GetCommentsRequest{PostUuid: postUuid, Offset: offset, Limit: limit})
	if err != nil {
		return result, fmt.Errorf("could not GetComments: %v", err)
	}

	comments := reply.GetComments()

	for _, commentPtr := range comments {
		result = append(result, ToGetCommentResult(commentPtr))
	}

	return result, nil
}

func (s *PostsGRPCService) GetCommentsStream(postUuid string, commentIds []int32) (<-chan (GetCommentResult), error) {
	var result chan (GetCommentResult) = make(chan GetCommentResult)
	var resultErr error
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetCommentsStream: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	stream, err := s.client.GetCommentsStream(ctx)
	if err != nil {
		return result, fmt.Errorf("could not GetCommentsStream: %v", err)
	}
	defer stream.CloseSend()

	if err != nil {
		return result, fmt.Errorf("s.client.GetCommentsStream failed: %v", err)
	}
	go func() {
		defer close(result)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				resultErr = fmt.Errorf("s.client.GetCommentsStream: stream.Recv failed: %v", err)
				return
			}
			result <- ToGetCommentResult(in)
		}
	}()
	for _, commentId := range commentIds {
		err := stream.Send(&GetCommentRequest{PostUuid: postUuid, Id: commentId})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetCommentsStream: stream.Send({postUuid: '%v', commentId: '%v'}) failed: %v", postUuid, commentId, err)
			return result, resultErr
		}
	}
	return result, resultErr
}

func ToGetPostsResult(post *GetPostReply) GetPostResult {
	return GetPostResult{
		Uuid:           post.GetUuid(),
		AuthorId:       int(post.GetAuthorId()),
		Text:           post.GetText(),
		PreviewText:    post.GetPreviewText(),
		Topic:          post.GetTopic(),
		State:          post.GetState(),
		CreateDate:     post.GetCreateDate().AsTime(),
		LastUpdateDate: post.GetLastUpdateDate().AsTime(),
	}
}

func ToGetCommentResult(comment *GetCommentReply) GetCommentResult {
	return GetCommentResult{
		Id:              int(comment.GetId()),
		Uuid:            comment.GetUuid(),
		AuthorId:        int(comment.GetAuthorId()),
		PostUuid:        comment.GetPostUuid(),
		LinkedCommentId: utils.Int32ToIntPtr(comment.GetLinkedCommentId()),
		Text:            comment.GetText(),
		State:           comment.GetState(),
		CreateDate:      comment.GetCreateDate().AsTime(),
		LastUpdateDate:  comment.GetLastUpdateDate().AsTime(),
	}
}
