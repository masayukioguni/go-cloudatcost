package cloudatcost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListTemplatesService_ListTemplates(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listtemplates.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("listtemplates.json"))
	})

	listtemplates, _, _ := mock.Client.ListTemplatesService.ListTemplates()
	wantListTemplateId := "26"
	if !reflect.DeepEqual(listtemplates[0].ID, wantListTemplateId) {
		t.Errorf("ListTemplatesService.ListTemplates returned %+v, want %+v", listtemplates[0].ID, wantListTemplateId)
	}
}

func TestListTemplatesService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/api/v1/listtemplates.php", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.ListTemplatesService.ListTemplates()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
		t.Errorf("ListTemplatesService.ListTemplates returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
	}

	if err == nil {
		t.Errorf("ListTemplatesService.ListTemplates returned %+v, want erespor message.", err)
	}

}
