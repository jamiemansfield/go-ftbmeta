package ftbmeta

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL   = "https://meta.ftb.neptunepowered.org/"
	defaultUserAgent = "go-ftbmeta"
)

// A client manages communication with The Neptune FTB Meta Service.
type Client struct {
	// HTTP client used to communicate.
	client *http.Client

	// Base URL for meta requests.
	BaseURL *url.URL

	// User Agent used when communicating.
	UserAgent string

	// Meta services.
	Packs *PackService
}

type service struct {
	client *Client
}

// NewClient returns a new modpacks.ch API client. If a nil client is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client: httpClient,
		BaseURL: baseURL,
		UserAgent: defaultUserAgent,
	}
	c.Packs = &PackService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided
// in urlStr, in which case it is resolved to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method string, urlStr string) (*http.Request, error) {
	// Resolve absolute URL
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Create the request
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response
// is JSON decoded and stored in the value pointed to by v.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
