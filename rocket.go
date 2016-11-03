package gorocket

import (
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"fmt"
)

type Rocket struct {
	Protocol string
	Host  string
	Port  string

	Debug bool

	auth  *authInfo
}

type authInfo struct {
	token string
	id    string
}

type statusResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func (r *Rocket) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", r.Protocol, r.Host, r.Port)
}

func (r *Rocket) doRequest(request *http.Request, responseBody interface{}) error {

	if r.auth != nil {
		request.Header.Set("X-Auth-Token", r.auth.token)
		request.Header.Set("X-User-Id", r.auth.id)
	}

	if r.Debug {
		log.Println(request)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	if r.Debug {
		log.Println(string(bodyBytes))
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Request error: " + response.Status)
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, responseBody)
}

