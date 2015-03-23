package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListTasksService_ListTasks(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listtasks.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("listtasks.json"))
	})

	listtasks, _, _ := mock.Client.ListTasksService.ListTasks()
	wantCid := "734103810"

	if !reflect.DeepEqual(listtasks[0].Cid, wantCid) {
		t.Errorf("ListTasksService.ListTasks returned %+v, want %+v", listtasks[0].Cid, wantCid)
	}
}

func TestListTasksService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listtasks.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.ListTasksService.ListTasks()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("ListTasksService.ListTasks returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("ListTasksService.ListTasks returned %+v, want erespor message.", err)
	}

}
