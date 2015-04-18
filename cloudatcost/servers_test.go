package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServersService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listservers.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("listservers.json"))
	})

	listservers, _, _ := mock.Client.ServersService.List()
	wantListServerId := "test id"

	if !reflect.DeepEqual(listservers[0].ID, wantListServerId) {
		t.Errorf("ServersService.List returned %+v, want %+v", listservers[0].ID, wantListServerId)
	}

}

func TestServersService_Rename(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/renameserver.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mock.ReadJSON("standard.json"))
	})

	res, _, _ := mock.Client.ServersService.Rename("1", "a")
	wantResult := "successful"

	if !reflect.DeepEqual(res.Result, wantResult) {
		t.Errorf("ServersService.List returned %+v, want %+v", res.Result, wantResult)
	}
}

func TestServersService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listservers.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.ServersService.List()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("ServersService.List returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("ServersService.List returned %+v, want erespor message.", err)
	}

}
