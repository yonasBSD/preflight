// Package dashboard is the client for the Preflight dashboard service
// (app.preflight.sh): credential storage and the HTTP API used by `preflight
// auth` and `preflight scan --publish`.
//
// The base URL is configurable via PREFLIGHT_API_URL so the whole flow can be
// exercised against a local server (e.g. http://localhost:8080) without a
// release.
package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DefaultAPIURL is the production dashboard origin.
const DefaultAPIURL = "https://app.preflight.sh"

// APIURL returns the dashboard base URL, honoring PREFLIGHT_API_URL.
func APIURL() string {
	if v := os.Getenv("PREFLIGHT_API_URL"); v != "" {
		return v
	}
	return DefaultAPIURL
}

// Credentials is the persisted CLI auth state.
type Credentials struct {
	Token  string `json:"token"`
	Email  string `json:"email"`
	APIURL string `json:"api_url"`
}

// credentialsPath returns ~/.preflight/credentials.json.
func credentialsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".preflight", "credentials.json"), nil
}

// LoadCredentials reads stored credentials, returning (nil, nil) when none exist.
func LoadCredentials() (*Credentials, error) {
	path, err := credentialsPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var c Credentials
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse credentials: %w", err)
	}
	return &c, nil
}

// Save writes credentials to ~/.preflight/credentials.json with 0600 perms.
func (c *Credentials) Save() error {
	path, err := credentialsPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// ClearCredentials removes any stored credentials (logout).
func ClearCredentials() error {
	path, err := credentialsPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Client talks to the dashboard API.
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// NewClient returns a Client pointed at APIURL().
func NewClient() *Client {
	return &Client{BaseURL: APIURL(), HTTP: &http.Client{Timeout: 30 * time.Second}}
}

// StartResponse is returned by StartAuth.
type StartResponse struct {
	DeviceCode string `json:"device_code"`
	UserCode   string `json:"user_code"`
	VerifyURL  string `json:"verify_url"`
	Interval   int    `json:"interval"`
}

// StartAuth begins the device-authorization flow.
func (c *Client) StartAuth() (*StartResponse, error) {
	resp, err := c.HTTP.Post(c.BaseURL+"/api/cli/auth/start", "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth start failed: %s", resp.Status)
	}
	var out StartResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PollStatus is the result of one Poll call.
type PollStatus struct {
	Status string // "pending", "approved", "expired"
	Token  string
}

// Poll checks once whether the device code has been approved.
func (c *Client) Poll(deviceCode string) (*PollStatus, error) {
	body, _ := json.Marshal(map[string]string{"device_code": deviceCode})
	resp, err := c.HTTP.Post(c.BaseURL+"/api/cli/auth/poll", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&out)
	return &PollStatus{Status: out.Status, Token: out.Token}, nil
}

// Whoami returns the email for a token, or an error if it is invalid.
func (c *Client) Whoami(token string) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, c.BaseURL+"/api/cli/whoami", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("token is not valid")
	}
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return "", fmt.Errorf("whoami failed: %s: %s", resp.Status, string(b))
	}
	var out struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.Email, nil
}
