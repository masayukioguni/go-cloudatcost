package cloudatcost

import (
	"net/http"
	"net/url"
)

type ConsoleService struct {
	client *Client
}

func (s *ConsoleService) Console(serverId string) (*ConsoleResponse, *http.Response, error) {
	u := "/api/v1/console.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	cr := new(ConsoleResponse)
	resp, err := s.client.Do(req, cr)
	if err != nil {
		return nil, resp, err
	}
	return cr, resp, err
}
