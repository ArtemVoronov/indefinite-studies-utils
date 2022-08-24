package auth

import (
	"bytes"
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	client         *http.Client
	baseURL        string
	verifyTokenURL string
}

type VerificationDTO struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

type VerificationResult struct {
	IsValid   bool
	IsExpired bool
}

func CreateAuthService(client *http.Client, baseUrl string) *AuthService {
	return &AuthService{
		client:         client,
		baseURL:        baseUrl,
		verifyTokenURL: baseUrl + "/api/v1/auth/verify-token",
	}
}

func (s *AuthService) Shutdown() error {
	return nil
}

func (s *AuthService) VerifyToken(token string) (*VerificationResult, error) {
	var result *VerificationResult

	body, err := json.Marshal(VerificationDTO{AccessToken: token})
	if err != nil {
		return result, fmt.Errorf("unable to verify token: %s", err)
	}

	req, err := http.NewRequest(http.MethodPost, s.verifyTokenURL, bytes.NewBuffer(body))
	if err != nil {
		return result, fmt.Errorf("unable to verify token: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := s.client.Do(req)
	if err != nil {
		return result, fmt.Errorf("unable to verify token: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("unable to verify token, response status code: %v", resp.StatusCode)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("unable to verify token: %s", err)
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return result, fmt.Errorf("unable to verify token: %s", err)
	}

	return result, nil
}

type AuthGRPCService struct {
	serverHost   string
	dialOptions  []grpc.DialOption
	connection   *grpc.ClientConn
	client       AuthServiceClient
	queryTimeout time.Duration
}

func CreateAuthGRPCService(serverHost string) *AuthGRPCService {
	// TODO: add TLS
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return &AuthGRPCService{
		serverHost:   serverHost,
		dialOptions:  opts,
		queryTimeout: 30 * time.Second,
	}
}

func (s *AuthGRPCService) connect() error {
	conn, err := grpc.Dial(s.serverHost, s.dialOptions...)
	if err != nil {
		return fmt.Errorf("unable to connect to '%v', error: %v", s.serverHost, err)
	}
	s.connection = conn
	s.client = NewAuthServiceClient(conn)
	return nil
}

func (s *AuthGRPCService) Shutdown() error {
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}

func (s *AuthGRPCService) ValidateCredentials(token string) (*VerificationResult, error) {
	var result *VerificationResult
	if s.connection == nil {
		err := s.connect()
		if err != nil {
			return result, fmt.Errorf("could not greet: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.queryTimeout)
	defer cancel()

	reply, err := s.client.VerifyToken(ctx, &VerifyTokenRequest{Token: token})
	if err != nil {
		return result, fmt.Errorf("could not greet: %v", err)
	}

	result = &VerificationResult{
		IsValid:   reply.GetIsValid(),
		IsExpired: reply.GetIsExpired(),
	}

	return result, nil
}
