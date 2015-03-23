package cloudatcost

import (
	"bytes"
	"net/http"
	"net/url"
)

type ConsoleService struct {
	client *Client
}

func (s *ConsoleService) Console(serverId string) (*ConsoleResponse, *http.Response, error) {
	urlStr := "/api/v1/console.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	u := s.client.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(parameters.Encode()))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	cr := new(ConsoleResponse)
	resp, err := s.client.Do(req, cr)
	if err != nil {
		return nil, resp, err
	}
	return cr, resp, err
}
