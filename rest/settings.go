package rest

import (
  "fmt"
  "net/http"
  "bytes"
  "github.com/skilld-labs/gorocket/api"
)

type settingResponse struct {
  Success  bool `json:"success"`
}

func (c *Client) Setting(s *api.Setting) error {
  var body string
  if s.Value == `true` || s.Value == `false` {
    body = fmt.Sprintf(`{"value": %s}`, s.Value)
  } else {
    body = fmt.Sprintf(`{"value": "%s"}`, s.Value)
  }
  request, _ := http.NewRequest("POST", c.getUrl() + "/api/v1/settings/" + s.Id, bytes.NewBufferString(body))

  return c.doRequest(request, &settingResponse{})
}
