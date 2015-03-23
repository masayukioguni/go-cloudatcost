package digitalocean

import (
	"fmt"
	"net/http"
)

// DropletsService Digital Ocean API docs: https://developers.digitalocean.com/#droplets
type DropletsService struct {
	client *Client
}

// CreateDropletsReque https://developers.digitalocean.com/#create-a-new-droplet
type CreateDropletRequest struct {
	Name              string   `json:"name,omitempty"`
	Region            string   `json:"region,omitempty"`
	Size              string   `json:"size,omitempty"`
	Image             string   `json:"image,omitempty"`
	SSHKeys           []string `json:"ssh_keys,omitempty"`
	Backups           bool     `json:"backups,omitempty"`
	IPv6              bool     `json:"ipv6,omitempty"`
	PrivateNetworking bool     `json:"private_networking,omitempty"`
	UserData          string   `json:"user_data,omitempty"`
}

// NetworkV4 The details of the network that are configured for the Droplet instance.
type NetworkV4 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// NetworkV6 The details of the network that are configured for the Droplet instance.
type NetworkV6 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   int    `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Networks The details of the network that are configured for the Droplet instance.
type Networks struct {
	V4s []NetworkV4 `json:"v4,omitempty"`
	V6s []NetworkV6 `json:"v6,omitempty"`
}

// Kernel The current kernel.
type Kernel struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// Droplet https://developers.digitalocean.com/#create-a-new-droplet
type Droplet struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Memory    int    `json:"memory,omitempty"`
	Vcpus     int    `json:"vcpus,omitempty"`
	Disk      int    `json:"disk,omitempty"`
	Locked    bool   `json:"locked,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Status    string `json:"status,omitempty"`

	BackupIds   []int    `json:"backup_ids,omitempty"`
	SnapshotIds []int    `json:"snapshot_ids,omitempty"`
	ActionIds   []string `json:"action_ids,omitempty"`
	Features    []string `json:"features,omitempty"`
	Region      Region   `json:"region,omitempty"`
	Size        string   `json:"size_slug,omitempty"`

	Networks Networks `json:"networks,omitempty"`
	Image    Image    `json:"image,omitempty"`
	Kernel   Kernel   `json:"kernel,omitempty"`
}

// DropletResponse https://developers.digitalocean.com/#create-a-new-droplet
type DropletResponse struct {
	Droplet Droplet `json:"droplet,omitempty"`
	//Links   Links   `json:"links,omitempty"`
}

// DropletsListResponse https://developers.digitalocean.com/#list-all-droplets
type DropletsListResponse struct {
	Droplets []Droplet `json:"droplets,omitempty"`
	//Links   Links   `json:"links,omitempty"`
}

// Create a new droplet
func (s *DropletsService) Create(droplet *CreateDropletRequest) (*Droplet, *http.Response, error) {

	u := "v2/droplets"

	droplet.IPv6 = true
	droplet.UserData = ""

	req, err := s.client.NewRequest("POST", u, droplet)
	if err != nil {
		return nil, nil, err
	}

	dr := new(DropletResponse)
	resp, err := s.client.Do(req, dr)

	if err != nil {
		return nil, resp, err
	}

	return &dr.Droplet, resp, err
}

// List List all droplets
func (s *DropletsService) List() ([]Droplet, *http.Response, error) {
	u := "v2/droplets"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	sr := new(DropletsListResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}

	return sr.Droplets, resp, err
}

// Destroy a droplet
func (s *DropletsService) Destroy(id int) (*http.Response, error) {
	u := fmt.Sprintf("v2/droplets/%v", id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Retrieve an existing Droplet by id
func (s *DropletsService) Get(id int) (*Droplet, *http.Response, error) {
	u := fmt.Sprintf("v2/droplets/%v", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	dr := new(DropletResponse)
	resp, err := s.client.Do(req, dr)

	if err != nil {
		return nil, resp, err
	}

	return &dr.Droplet, resp, err

}
