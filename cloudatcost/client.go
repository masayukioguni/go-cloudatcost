package cloudatcost

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
	defaultBaseURL = "https://panel.cloudatcost.com/api/v1/"
	userAgent      = "a"
)

// Option Optional parameters
type Option struct {
	Login string
	Key   string
}

// ErrorStatus https://github.com/cloudatcost/api Error:
type ErrorStatus struct {
	Status           string `json:"status,omitempty"`
	Time             string `json:"time,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// A Client manages communication with the CloudAtCost API.
type Client struct {
	Option                 *Option
	client                 *http.Client
	BaseURL                *url.URL
	UserAgent              string
	ListServersService     *ListServersService
	ListTemplatesService   *ListTemplatesService
	ListTasksService       *ListTasksService
	PowerOperationsService *PowerOperationsService
	ConsoleService         *ConsoleService
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response    *http.Response
	ErrorStatus ErrorStatus
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorStatus.Status, r.ErrorStatus.Time, r.ErrorStatus.Error, r.ErrorStatus.ErrorDescription)
}

// NewClient returns a new CloudAtCost API client.
func NewClient(option *Option) (*Client, error) {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Option: option}
	c.ListServersService = &ListServersService{client: c}
	c.ListTemplatesService = &ListTemplatesService{client: c}
	c.ListTasksService = &ListTasksService{client: c}
	c.PowerOperationsService = &PowerOperationsService{client: c}
	c.ConsoleService = &ConsoleService{client: c}

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
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) NewFormRequest(method, urlStr string, values url.Values) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf *bytes.Buffer
	if values != nil {
		buf = bytes.NewBufferString(values.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
