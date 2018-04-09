package kongo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	version   = "v0"
	userAgent = "kongo/" + version
	mediaType = "application/json"
)

type (
	// Kongo manages communication with Kong Admin API.
	Kongo struct {

		// HTTP client used to communicate with the Kong Admin API.
		client *http.Client

		// Kong server base URL.
		BaseURL *url.URL

		// User agent for client
		UserAgent string

		// Node api service
		Node Node
	}

	// An ErrorResponse report the error caused by and API request
	ErrorResponse struct {
		// HTTP response that caused this error
		Response *http.Response

		// Error message based on http status code
		Message string `json:"message, omitempty"`
	}

	// Custom time struct for json parsing
	Time struct {
		time.Time
	}
)

// NewClient returns a new Kongo API client.
func NewClient(client *http.Client, baseURL *url.URL) (*Kongo, error) {
	if client == nil {
		client = http.DefaultClient
	}

	if baseURL == nil {
		return nil, errors.New("Empty URL is not allowed")
	}

	k := &Kongo{client: client, BaseURL: baseURL, UserAgent: userAgent}
	k.Node = &NodeService{k}

	return k, nil
}

// New returns a new Kongo API client.
func New(client *http.Client, baseURL string) (*Kongo, error) {
	if baseURL == "" {
		return nil, errors.New("Empty URL is not allowed")
	}

	parsedURL, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	return NewClient(client, parsedURL)
}

// NewRequest creates an API requrest. A relative URL can be provided in res URL instance.
func (k *Kongo) NewRequest(ctx context.Context, method string, res *url.URL) (*http.Request, error) {
	url := k.BaseURL.ResolveReference(res)

	req, err := http.NewRequest(method, url.String(), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request and returns the API response. If the HTTP response is in the 2xx range,
// unmarshal the response body into value.
func (k *Kongo) Do(req *http.Request, value interface{}) (*http.Response, error) {
	res, err := k.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = k.checkResponse(res)

	if err != nil {
		return res, err
	}

	if value == nil {
		return res, nil
	}

	err = json.NewDecoder(res.Body).Decode(value)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// checkResponse checks the API response for errors and returns them if present.
func (k *Kongo) checkResponse(res *http.Response) error {
	if c := res.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: res}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	err = json.Unmarshal(data, errorResponse)

	if err != nil {
		errorResponse.Message = string(data)
	}

	return errorResponse
}

// Error retrieves the error message of Error Response
func (e *ErrorResponse) Error() string {
	if e.Message == "" {
		e.Message = "Request error"
	}

	return fmt.Sprintf(
		"%d %s",
		e.Response.StatusCode,
		e.Message,
	)
}

func (t *Time) UnmarshalJSON(value []byte) (err error) {
	v := strings.Trim(string(value), "\"")

	if v == "" {
		t.Time = time.Time{}

		return
	}

	timestamp, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return
	}

	t.Time = time.Unix(timestamp, 0)

	return
}
