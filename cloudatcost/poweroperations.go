package cloudatcost

import (
	"bytes"
	"net/http"
	"net/url"
)

type PowerOperationsService struct {
	client *Client
}

func (s *PowerOperationsService) Action(serverId, action string) (*PowerOperationResponse, *http.Response, error) {
	urlStr := "/api/v1/powerop.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)
	parameters.Add("action", action)

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	u := s.client.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(parameters.Encode()))
	if err != nil {
		return nil, nil, err
	}
	/*
		req, err := s.client.NewRequest("POST", u, bytes.NewBufferString(parameters.Encode()))
		if err != nil {
			return nil, nil, err
		}*/

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	por := new(PowerOperationResponse)
	resp, err := s.client.Do(req, por)
	if err != nil {
		return nil, resp, err
	}
	return por, resp, err
}

func (s *PowerOperationsService) PowerOn(serverId string) (*PowerOperationResponse, *http.Response, error) {
	return s.Action(serverId, "poweron")
}
func (s *PowerOperationsService) PowerOff(serverId string) (*PowerOperationResponse, *http.Response, error) {
	return s.Action(serverId, "poweroff")
}
func (s *PowerOperationsService) Reset(serverId string) (*PowerOperationResponse, *http.Response, error) {
	return s.Action(serverId, "reset")
}
