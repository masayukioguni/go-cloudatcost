package cloudatcost

import (
	"fmt"
	"net/http"
	"net/url"
)

type ServersService struct {
	client *Client
}

func (s *ServersService) List() ([]ListServer, *http.Response, error) {
	u := fmt.Sprintf("/api/v1/listservers.php?key=%s&login=%s", s.client.Option.Key, s.client.Option.Login)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	lsr := new(ListServersResponse)
	resp, err := s.client.Do(req, lsr)
	if err != nil {
		return nil, resp, err
	}
	return lsr.Data, resp, err
}

func (s *ServersService) Rename(serverId, name string) (*StandardResponse, *http.Response, error) {
	u := "/api/v1/renameserver.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)
	parameters.Add("name", name)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	sr := new(StandardResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}
	return sr, resp, err
}
