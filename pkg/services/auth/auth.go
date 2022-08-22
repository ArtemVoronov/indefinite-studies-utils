package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {
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

func CreateAuthService(client *http.Client, baseUrl string) *Service {
	return &Service{
		client:         client,
		baseURL:        baseUrl,
		verifyTokenURL: baseUrl + "/api/v1/auth/verify-token",
	}
}

func (s *Service) VerifyToken(token string) (*VerificationResult, error) {
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
