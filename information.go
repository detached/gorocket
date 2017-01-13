package gorocket

import (
	"net/http"
)

type versionResponse struct {
	Info Versions `json:"info"`
}

type Versions struct {
	Version string `json:"version"`
}

// Get version information from the server.
// This function does not need a logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (r *Rocket) GetVersions() (*Versions, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/info", nil)

	response := new(versionResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Info, nil;
}