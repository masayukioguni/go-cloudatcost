package cloudatcost

import (
	"fmt"
	"net/http"
	"net/url"
)

type CloudProService struct {
	client *Client
}

func (s *CloudProService) Resources() (*CloudProResourcesData, *http.Response, error) {
	u := fmt.Sprintf("/api/v1/cloudpro/resources.php?key=%s&login=%s", s.client.Option.Key, s.client.Option.Login)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rr := new(CloudProResourcesResponse)
	resp, err := s.client.Do(req, rr)
	if err != nil {
		return nil, resp, err
	}
	return &rr.Data, resp, err
}

func (s *CloudProService) Create(opt *CreateServerOptions) (*CloudProServerResponse, *http.Response, error) {
	u := "/api/v1/cloudpro/build.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("cpu", opt.Cpu)
	parameters.Add("os", opt.OS)
	parameters.Add("ram", opt.Ram)
	parameters.Add("storage", opt.Storage)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	sr := new(CloudProServerResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}
	return sr, resp, err
}

func (s *CloudProService) Delete(sid string) (*CloudProServerResponse, *http.Response, error) {
	u := "/api/v1/cloudpro/delete.php"

	parameters := url.Values{}
	parameters.Add("key", s.client.Option.Key)
	parameters.Add("login", s.client.Option.Login)
	parameters.Add("sid", sid)

	req, err := s.client.NewFormRequest("POST", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	sr := new(CloudProServerResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}
	return sr, resp, err
}
