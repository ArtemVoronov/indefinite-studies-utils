package posts

import (
	context "context"
	"fmt"
	"io"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type GetPostResult struct {
	Id             int
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
	AuthorId        int
	PostId          int
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

func (s *PostsGRPCService) GetPost(postId int32) (*GetPostResult, error) {
	var result *GetPostResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetPost: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetPost(ctx, &GetPostRequest{Id: postId})
	if err != nil {
		return result, fmt.Errorf("could not GetPost: %v", err)
	}

	result = &GetPostResult{
		Id:             int(reply.GetId()),
		AuthorId:       int(reply.GetAuthorId()),
		Text:           reply.GetText(),
		PreviewText:    reply.GetPreviewText(),
		Topic:          reply.GetTopic(),
		State:          reply.GetState(),
		CreateDate:     reply.GetCreateDate().AsTime(),
		LastUpdateDate: reply.GetLastUpdateDate().AsTime(),
	}

	return result, nil
}

func (s *PostsGRPCService) GetPosts(offset int32, limit int32, userIds []int32) ([]GetPostResult, error) {
	var result []GetPostResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetPosts: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetPosts(ctx, &GetPostsRequest{Offset: offset, Limit: limit, Ids: userIds})
	if err != nil {
		return result, fmt.Errorf("could not GetPosts: %v", err)
	}

	posts := reply.GetPosts()

	for _, postPtr := range posts {
		result = append(result, GetPostResult{
			Id:             int(postPtr.GetId()),
			AuthorId:       int(postPtr.GetAuthorId()),
			Text:           postPtr.GetText(),
			PreviewText:    postPtr.GetPreviewText(),
			Topic:          postPtr.GetTopic(),
			State:          postPtr.GetState(),
			CreateDate:     postPtr.GetCreateDate().AsTime(),
			LastUpdateDate: postPtr.GetLastUpdateDate().AsTime(),
		})
	}

	return result, nil
}

func (s *PostsGRPCService) GetPostsStream(userIds []int32) (<-chan (GetPostResult), error) {
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
			result <- GetPostResult{
				Id:             int(in.GetId()),
				AuthorId:       int(in.GetAuthorId()),
				Text:           in.GetText(),
				PreviewText:    in.GetPreviewText(),
				Topic:          in.GetTopic(),
				State:          in.GetState(),
				CreateDate:     in.GetCreateDate().AsTime(),
				LastUpdateDate: in.GetLastUpdateDate().AsTime(),
			}
		}
	}()
	for _, userId := range userIds {
		err := stream.Send(&GetPostRequest{Id: userId})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetPostsStream: stream.Send(%v) failed: %v", userId, err)
			return result, resultErr
		}
	}
	return result, resultErr
}

func (s *PostsGRPCService) GetComment(postId int32) (*GetCommentResult, error) {
	var result *GetCommentResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetComment: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetComment(ctx, &GetCommentRequest{Id: postId})
	if err != nil {
		return result, fmt.Errorf("could not GetComment: %v", err)
	}

	result = &GetCommentResult{
		Id:              int(reply.GetId()),
		AuthorId:        int(reply.GetAuthorId()),
		PostId:          int(reply.GetPostId()),
		LinkedCommentId: parseNillable(reply.GetLinkedCommentId()),
		Text:            reply.GetText(),
		State:           reply.GetState(),
		CreateDate:      reply.GetCreateDate().AsTime(),
		LastUpdateDate:  reply.GetLastUpdateDate().AsTime(),
	}

	return result, nil
}

func (s *PostsGRPCService) GetComments(postId int32, offset int32, limit int32) ([]GetCommentResult, error) {
	var result []GetCommentResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not GetComments: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.GetComments(ctx, &GetCommentsRequest{PostId: postId, Offset: offset, Limit: limit})
	if err != nil {
		return result, fmt.Errorf("could not GetComments: %v", err)
	}

	comments := reply.GetComments()

	for _, commentPtr := range comments {
		result = append(result, GetCommentResult{
			Id:              int(commentPtr.GetId()),
			AuthorId:        int(commentPtr.GetAuthorId()),
			PostId:          int(commentPtr.GetPostId()),
			LinkedCommentId: parseNillable(commentPtr.GetLinkedCommentId()),
			Text:            commentPtr.GetText(),
			State:           commentPtr.GetState(),
			CreateDate:      commentPtr.GetCreateDate().AsTime(),
			LastUpdateDate:  commentPtr.GetLastUpdateDate().AsTime(),
		})
	}

	return result, nil
}

func (s *PostsGRPCService) GetCommentsStream(userIds []int32) (<-chan (GetCommentResult), error) {
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
			result <- GetCommentResult{
				Id:              int(in.GetId()),
				AuthorId:        int(in.GetAuthorId()),
				PostId:          int(in.GetPostId()),
				LinkedCommentId: parseNillable(in.GetLinkedCommentId()),
				Text:            in.GetText(),
				State:           in.GetState(),
				CreateDate:      in.GetCreateDate().AsTime(),
				LastUpdateDate:  in.GetLastUpdateDate().AsTime(),
			}
		}
	}()
	for _, userId := range userIds {
		err := stream.Send(&GetCommentRequest{Id: userId})
		if err != nil {
			resultErr = fmt.Errorf("s.client.GetCommentsStream: stream.Send(%v) failed: %v", userId, err)
			return result, resultErr
		}
	}
	return result, resultErr
}

func parseNillable(val int32) *int {
	if val == 0 {
		return nil
	}
	result := int(val)
	return &result
}
