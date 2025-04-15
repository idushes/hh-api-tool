package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	// HeadHunter OAuth token endpoint
	HHTokenEndpoint = "https://api.hh.ru/oauth/token"
)

// TokenResponse represents the response from the HeadHunter OAuth token endpoint
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// ErrorResponse represents an error from the HeadHunter API
type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// HeadHunterAuthClient provides methods for authenticating with the HeadHunter API
type HeadHunterAuthClient struct {
	httpClient *http.Client
}

// NewHeadHunterAuthClient creates a new HeadHunterAuthClient
func NewHeadHunterAuthClient() *HeadHunterAuthClient {
	return &HeadHunterAuthClient{
		httpClient: &http.Client{},
	}
}

// GetClientCredentialsToken obtains a token using client credentials
func (c *HeadHunterAuthClient) GetClientCredentialsToken(clientID, clientSecret string) (*TokenResponse, error) {
	// Prepare form data
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	// Create request
	req, err := http.NewRequest("POST", HHTokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check for error responses
	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("error decoding error response: %w", err)
		}
		return nil, fmt.Errorf("API error: %s - %s", errResp.Error, errResp.ErrorDescription)
	}

	// Decode response
	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("error decoding token response: %w", err)
	}

	return &tokenResp, nil
}
