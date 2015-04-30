package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCloudProService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/cloudpro/resources.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("cloudpro/resources.json"))
	})

	res, _, _ := mock.Client.CloudProService.Resources()
	want := "8"

	if !reflect.DeepEqual(res.Total[0].CPU, want) {
		t.Errorf("CloudProService.Resources returned %+v, want %+v", res.Total[0].CPU, want)
	}
}

func TestCloudProService_Create(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/cloudpro/build.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mock.ReadJSON("cloudpro/create.json"))
	})

	opt := CreateServerOptions{}

	res, _, _ := mock.Client.CloudProService.Create(&opt)
	want := "build"

	if !reflect.DeepEqual(res.Action, want) {
		t.Errorf("CloudProService.Create returned %+v, want %+v", res.Action, want)
	}
}

func TestCloudProService_Delete(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/cloudpro/delete.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, mock.ReadJSON("cloudpro/delete.json"))
	})

	res, _, _ := mock.Client.CloudProService.Delete("1")
	want := "delete"

	if !reflect.DeepEqual(res.Action, want) {
		t.Errorf("CloudProService.Delete returned %+v, want %+v", res.Action, want)
	}
}

func TestCloudProService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/cloudpro/resources.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.CloudProService.Resources()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("CloudProService.Resources returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("CloudProService.Resources returned %+v, want erespor message.", err)
	}

}
