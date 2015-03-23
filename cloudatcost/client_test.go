package digitalocean

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, _ := NewClient(&Option{APIKey: "Test"})

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestDo_httpError(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := mock.Client.NewRequest("GET", "/", nil)
	_, err := mock.Client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"id":"error id.","message":"error message."}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
		ErrorStatus: ErrorStatus{
			ID:      "error id.",
			Message: "error message.",
		},
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestCheckResponse_noBody(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}
	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestErrorResponse_Error(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	err := &ErrorResponse{
		Response: res,
	}
	if err.Error() == "" {
		t.Errorf("Expected non-empty Error.Error()")
	}
}
