package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPowerOperationsService_PowerOperations(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/powerop.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mock.ReadJSON("powerop.json"))
	})

	po, _, _ := mock.Client.PowerOperationsService.PowerOn("1")
	want := "poweron"
	if !reflect.DeepEqual(po.Action, want) {
		t.Errorf("PowerOperationsService.PowerOn returned %+v, want %+v", po.Action, want)
	}
}

func TestPowerOperationsService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/powerop.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.PowerOperationsService.PowerOn("1")

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("PowerOperationsService.PowerOn returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("PowerOperationsService.PowerOn returned %+v, want erespor message.", err)
	}

}
