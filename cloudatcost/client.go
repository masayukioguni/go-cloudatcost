package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// const
const (
	LibraryVersion = "0.1"
	defaultBaseURL = "https://api.digitalocean.com/"
	userAgent      = "go-tugboat/" + LibraryVersion
)

// Meta https://developers.digitalocean.com/#meta
type Meta struct {
	Total int `json:"total,omitempty"`
}

// Option Optional parameters
type Option struct {
	APIKey string
}

// ErrorStatus https://developers.digitalocean.com/#statuses
type ErrorStatus struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// A Client manages communication with the Digital Ocean API.
type Client struct {
	Option          *Option
	client          *http.Client
	BaseURL         *url.URL
	UserAgent       string
	SizesService    *SizesService
	RegionsService  *RegionsService
	DropletsService *DropletsService
	ImagesService   *ImagesService
	SSHKeysService  *SSHKeysService
	AccountService  *AccountService
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response    *http.Response
	ErrorStatus ErrorStatus
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorStatus.ID, r.ErrorStatus.Message)
}

// NewClient returns a new Digital Ocean API client.
func NewClient(option *Option) (*Client, error) {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Option: option}
	c.SizesService = &SizesService{client: c}
	c.RegionsService = &RegionsService{client: c}
	c.DropletsService = &DropletsService{client: c}
	c.ImagesService = &ImagesService{client: c, Page: 1, PerPage: 50}
	c.SSHKeysService = &SSHKeysService{client: c}
	c.AccountService = &AccountService{client: c}

	return c, nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	authHeader := "Bearer " + c.Option.APIKey

	req.Header.Set("Authorization", authHeader)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err

}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse.ErrorStatus)
	}

	return errorResponse
}
