package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDnsService_ModifyReverseDns(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/rdns.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mock.ReadJSON("standard.json"))
	})

	res, _, _ := mock.Client.DnsService.ModifyReverseDns("1", "a")
	wantResult := "successful"

	if !reflect.DeepEqual(res.Result, wantResult) {
		t.Errorf("DnsService.ModifyReverseDns returned %+v, want %+v", res.Result, wantResult)
	}
}

func TestDnsService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/rdns.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.DnsService.ModifyReverseDns("1", "a")

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("DnsService.ModifyReverseDns returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("DnsService.ModifyReverseDns returned %+v, want erespor message.", err)
	}

}
