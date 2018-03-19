package modica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.modicagroup.com/rest/gateway"
	userAgent      = "go-modica"
)

const (
	methodPost = "POST"
	methodGet  = "GET"
)

const (
	mediaTypeV1 = "application/vnd.modica.gateway.v1+json"
)

// Client enables talking to Modica's API
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Modica API.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the Modica API.
	UserAgent string

	// Authentication Details
	clientID     string
	clientSecret string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Modica API.
	MobileGateway *MobileGatewayService
}

type service struct {
	client *Client
}

// NewClient returns a new Modica API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(clientID string, clientSecret string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		BaseURL:      baseURL,
		client:       httpClient,
		clientID:     clientID,
		clientSecret: clientSecret,
		UserAgent:    userAgent,
	}
	c.common.client = c

	// Services
	c.MobileGateway = (*MobileGatewayService)(&c.common)

	return c
}

func (c *Client) newRequest(method string, urlPath string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	uri, err := c.BaseURL.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	// Configure Headers
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Accept", mediaTypeV1)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
