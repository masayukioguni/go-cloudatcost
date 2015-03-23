package cloudatcost

import (
	"fmt"
	"net/http"
)

type ListTasksService struct {
	client *Client
}

func (s *ListTasksService) ListTasks() ([]ListTask, *http.Response, error) {
	u := fmt.Sprintf("/api/v1/listtasks.php?key=%s&login=%s", s.client.Option.Key, s.client.Option.Login)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ltr := new(ListTasksResponse)
	resp, err := s.client.Do(req, ltr)
	if err != nil {
		return nil, resp, err
	}
	return ltr.Data, resp, err
}
