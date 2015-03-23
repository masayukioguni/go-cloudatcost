package cloudatcost

import (
	"fmt"
	"net/http"
)

type ListServersService struct {
	client *Client
}

func (s *ListServersService) ListServers() ([]ListServer, *http.Response, error) {
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
