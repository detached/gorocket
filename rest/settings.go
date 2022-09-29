package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/davidferlay/gorocket/api"
)

type settingResponse struct {
	Success bool `json:"success"`
}

func (c *Client) Setting(s *api.Setting) error {
	body, err := json.Marshal(struct {
		Value string `json:"value"`
	}{
		Value: s.Value,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/settings/"+s.Id, bytes.NewReader(body))

	return c.doRequest(request, &settingResponse{})
}
