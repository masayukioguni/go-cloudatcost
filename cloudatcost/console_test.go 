package cloudatcost

import (
  "fmt"
  "net/http"
  "reflect"
  "testing"
)

func TestConsoleService_Console(t *testing.T) {
  mock := NewMockClient()
  defer mock.Close()

  mock.Mux.HandleFunc("/api/v1/console.php", func(w http.ResponseWriter, r *http.Request) {
    testMethod(t, r, "POST")
    fmt.Fprint(w, mock.ReadJSON("console.json"))
  })

  res, _, _ := mock.Client.ConsoleService.Console("1")
  want := "http://panel.cloudatcost.com:12345/console.html?servername=123456&hostname=1.1.1.1&sshkey=123456&sha1hash=aBcDeFgG"
  if !reflect.DeepEqual(res.Console, want) {
    t.Errorf("ConsoleService.Console returned %+v, want %+v", res.Console, want)
  }
}

func TestConsoleService_Unauthorized(t *testing.T) {
  mock := NewMockClient()
  defer mock.Close()

  mock.Mux.HandleFunc("/api/v1/console.php", func(w http.ResponseWriter, r *http.Request) {
    testMethod(t, r, "POST")
    w.WriteHeader(http.StatusPreconditionFailed)
    fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
  })

  _, hr, err := mock.Client.ConsoleService.Console("1")

  if !reflect.DeepEqual(hr.StatusCode, http.StatusPreconditionFailed) {
    t.Errorf("ConsoleService.Console returned %+v, want %+v", hr.StatusCode, http.StatusPreconditionFailed)
  }

  if err == nil {
    t.Errorf("ConsoleService.Console returned %+v, want erespor message.", err)
  }

}
