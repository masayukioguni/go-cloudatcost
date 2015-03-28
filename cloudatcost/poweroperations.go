package cloudatcost

import (
	"net/http"
	"net/url"
)

type PowerOperationsService struct {
	client *Client
}

func (s *PowerOperationsService) Action(serverId, action string) (*PowerOperationResponse, *http.Response, error) {
	u := "/api/v1/powerop.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)
	parameters.Add("action", action)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

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
