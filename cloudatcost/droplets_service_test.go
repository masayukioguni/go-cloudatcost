package digitalocean

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDropletsService_Create(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	input := &CreateDropletRequest{
		Name:              "example.com",
		Size:              "512mb",
		Image:             "ubuntu-14-04-x64",
		Region:            "nyc3",
		SSHKeys:           nil,
		Backups:           false,
		PrivateNetworking: false}

	mock.Mux.HandleFunc("/v2/droplets", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateDropletRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, mock.ReadJSON("create_droplet.json"))
	})

	droplet, _, _ := mock.Client.DropletsService.Create(input)

	if !reflect.DeepEqual(droplet.ID, 3164494) {
		t.Errorf("DropletsService.Create returned %+v, want %+v", droplet.ID, 3164494)
	}
}

func TestDropletsService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/droplets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("list_all_droplets.json"))
	})

	droplets, _, _ := mock.Client.DropletsService.List()

	if !reflect.DeepEqual(droplets[0].ID, 3164444) {
		t.Errorf("DropletsService.Create returned %+v, want %+v", droplets[0].ID, 3164444)
	}
}

func TestDropletsService_Destroy(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/droplets/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	resp, _ := mock.Client.DropletsService.Destroy(12345)
	if !reflect.DeepEqual(resp.StatusCode, 200) {
		t.Errorf("DropletsService.Destory returned %+v, want %+v", resp, 200)
	}
}

func TestDropletsService_Get(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/droplets/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("retrieve_an_existing_droplet_by_id.json"))
	})

	droplet, _, _ := mock.Client.DropletsService.Get(12345)

	if !reflect.DeepEqual(droplet.ID, 3164494) {
		t.Errorf("DropletsService.Get returned %+v, want %+v", droplet.ID, 3164494)
	}
}
