package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// HeadHunter API base URL
	HHAPIBaseURL = "https://api.hh.ru"
)

// UserResponse represents the response from the HeadHunter /me endpoint
type UserResponse struct {
	ID              string      `json:"id"`
	Email           string      `json:"email"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	MiddleName      string      `json:"middle_name"`
	IsEmployer      bool        `json:"is_employer"`
	IsApplicant     bool        `json:"is_applicant"`
	IsAdmin         bool        `json:"is_admin"`
	Phone           string      `json:"phone"`
	PhotoID         string      `json:"photo_id"`
	CountryID       string      `json:"country_id"`
	AuthType        string      `json:"auth_type"`
	HasPassword     bool        `json:"has_password"`
	NegotiationsURL string      `json:"negotiations_url"`
	ResumesURL      string      `json:"resumes_url"`
	Employer        interface{} `json:"employer,omitempty"`
	Manager         interface{} `json:"manager,omitempty"`
	ProfileURL      string      `json:"profile_url"`
	FullName        string      `json:"full_name"`
}

// HeadHunterUserClient provides methods for working with user data in the HeadHunter API
type HeadHunterUserClient struct {
	httpClient *http.Client
}

// NewHeadHunterUserClient creates a new HeadHunterUserClient
func NewHeadHunterUserClient() *HeadHunterUserClient {
	return &HeadHunterUserClient{
		httpClient: &http.Client{},
	}
}

// GetCurrentUser retrieves information about the current authenticated user
func (c *HeadHunterUserClient) GetCurrentUser(accessToken string) (*UserResponse, error) {
	// Create request
	req, err := http.NewRequest("GET", HHAPIBaseURL+"/me", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("User-Agent", "HeadHunter API Tool (idushes/hh-api-tool)")
	req.Header.Set("Accept", "application/json")

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
			return nil, fmt.Errorf("error response with status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("API error: %s - %s", errResp.Error, errResp.ErrorDescription)
	}

	// Decode response
	var userResp UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("error decoding user response: %w", err)
	}

	return &userResp, nil
}
