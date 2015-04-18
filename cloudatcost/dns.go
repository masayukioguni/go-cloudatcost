package cloudatcost

import (
	"net/http"
	"net/url"
)

type DnsService struct {
	client *Client
}

func (s *DnsService) ModifyReverseDns(serverId, hostname string) (*StandardResponse, *http.Response, error) {
	u := "/api/v1/rdns.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", serverId)
	parameters.Add("hostname", hostname)

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
