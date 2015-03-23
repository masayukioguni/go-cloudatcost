package digitalocean

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRegionsService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/regions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("regions.json"))
	})

	regions, _, _ := mock.Client.RegionsService.List()

	wantRegion := Region{
		Name: string("San Francisco 1"),
		Slug: string("sfo1"),
		Sizes: []string{"32gb",
			"16gb",
			"2gb",
			"1gb",
			"4gb",
			"8gb",
			"512mb",
			"64gb",
			"48gb"},
		Features: []string{"virtio",
			"backups",
			"metadata"},
		Available: true,
	}

	if !reflect.DeepEqual(regions[0], wantRegion) {
		t.Errorf("RegionsService.List returned %+v, want %+v", regions[0], wantRegion)
	}

}

func TestRegionsService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/regions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.RegionsService.List()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusUnauthorized) {
		t.Errorf("RegionsService.List returned %+v, want %+v", hr.StatusCode, http.StatusUnauthorized)
	}

	if err == nil {
		t.Errorf("RegionsService.List returned %+v, want erespor message.", err)
	}

}
