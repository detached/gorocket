package gorocket

import (
	"errors"
	"net/http"
)

type versionResponse struct{
	statusResponse
	Versions Versions `json:"versions"`
}

type Versions struct {
	Api string `json:"api"`
	RocketChat string `json:"rocketchat"`
}

func (r *Rocket) GetVersions() (*Versions, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/version", nil)

	response := new(versionResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	if response.Status == "success" {
		return &response.Versions, nil
	} else {
		return nil, errors.New("Response status: " + response.Status)
	}
}