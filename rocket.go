package gorocket

import (
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
	"bytes"
	"log"
	"net/url"
)

type Rocket struct {
	Url       string
	Debug     bool

	auth *authInfo
}

type authInfo struct {
	token string
	id    string
}

func (r *Rocket) doRequest(request *http.Request, responseBody interface{}) error {

	if r.auth != nil {
		request.Header.Set("X-Auth-Token", r.auth.token)
		request.Header.Set("X-User-Id", r.auth.id)
	}

	if r.Debug {
		log.Println(request)
	}

	response, requestError := http.DefaultClient.Do(request)

	if requestError != nil {
		return requestError
	}

	defer response.Body.Close()
	bodyBytes, readError := ioutil.ReadAll(response.Body)

	if r.Debug {
		log.Println(string(bodyBytes))
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Request error: " + response.Status)
	}

	if readError != nil {
		return readError
	}

	return json.Unmarshal(bodyBytes, responseBody)
}

func (r *Rocket) Logon(username string, password string) error {
	data := url.Values{"user": {username}, "password": {password}}
	request, _ := http.NewRequest("POST", r.Url + "/api/login", bytes.NewBufferString(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := new(logonResponse)
	err := r.doRequest(request, response)

	if err != nil {
		return err
	}

	r.auth = &authInfo{id:response.Data.UserId, token:response.Data.Token}
	return nil
}

func (r *Rocket) GetPublicRooms() ([]Room, error) {
	request, _ := http.NewRequest("GET", r.Url + "/api/publicRooms", nil)
	request.Header.Set("Accept", "application/json")

	response := new(roomsResponse)
	err := r.doRequest(request, response)

	if err != nil {
		return nil, err
	}

	if response.Status == "success" {
		return response.Rooms, nil
	} else {
		return nil, errors.New("Response status: " + response.Status)
	}
}

func (r *Rocket) Send(room *Room, msg string) error {

	b, _ := json.Marshal(&message{Text:msg})

	request, _ := http.NewRequest("POST", r.Url + "/api/rooms/" + room.Id + "/send", bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response := new(statusResponse)
	err := r.doRequest(request, response)

	if err != nil {
		return err
	}

	if response.Status != "success" {
		return errors.New("Response status: " + response.Status)
	}

	return nil
}