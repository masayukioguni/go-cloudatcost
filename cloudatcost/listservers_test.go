package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListServersService_ListServers(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listservers.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("listservers.json"))
	})

	listservers, _, _ := mock.Client.ListServersService.ListServers()
	wantListServerId := "test id"

	if !reflect.DeepEqual(listservers[0].ID, wantListServerId) {
		t.Errorf("ListServersService.ListServers returned %+v, want %+v", listservers[0].ID, wantListServerId)
	}
}

func TestListServersService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listservers.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.ListServersService.ListServers()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("ListServersService.ListServers returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("ListServersService.ListServers returned %+v, want erespor message.", err)
	}

}
