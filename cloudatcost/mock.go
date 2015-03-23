package digitalocean

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
)

// A MockClient manages communication with the Digital Ocean API mock.
type MockClient struct {
	Mux         *http.ServeMux
	Client      *Client
	Server      *httptest.Server
	FixturesDir string
}

// NewMockClient returns a new Digital Ocean API client mock.
func NewMockClient() *MockClient {
	client, _ := NewClient(&Option{APIKey: "test"})

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url

	cli := &MockClient{Mux: mux,
		Server:      server,
		Client:      client,
		FixturesDir: "./test-fixtures"}

	return cli
}

// Close Close Mocl
func (r *MockClient) Close() {
	r.Server.Close()
}

// ReadJSON Read json from file
func (r *MockClient) ReadJSON(filename string) string {
	text, _ := ioutil.ReadFile(filepath.Join(r.FixturesDir, filename))
	return string(text)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %s, want %s", header, got, want)
	}
}
