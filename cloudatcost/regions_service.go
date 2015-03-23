package digitalocean

import (
	"net/http"
)

// RegionsService Digital Ocean API docs: https://developers.digitalocean.com/#regions
type RegionsService struct {
	client *Client
}

// Region https://developers.digitalocean.com/#regions
type Region struct {
	Slug      string   `json:"slug,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}

// RegionsResponse https://developers.digitalocean.com/#list-all-regions
type RegionsResponse struct {
	Regions []Region `json:"regions,omitempty"`
	Meta    Meta     `json:"meta,omitempty"`
}

// List https://developers.digitalocean.com/#list-all-regions
func (s *RegionsService) List() ([]Region, *http.Response, error) {
	u := "v2/regions"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rr := new(RegionsResponse)
	resp, err := s.client.Do(req, rr)
	if err != nil {
		return nil, resp, err
	}
	return rr.Regions, resp, err
}
