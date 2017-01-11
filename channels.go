package gorocket

import (
	"fmt"
	"net/http"
	"bytes"
)

type channelResponse struct {
	Success  bool `json:"success"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Id           string `json:"_id"`
	Name         string `json:"name"`
	MessageCount int `json:"msgs"`
	UserNames    []string `json:"usernames"`

	User         User `json:"u"`

	ReadOnly     bool `json:"ro"`
	Timestamp    string `json:"ts"`
	T            string `json:"t"`
	UpdatedAt    string `json:"_updatedAt"`
	SysMes       bool `json:"sysMes"`
}

func (r *Rocket) GetPublicChannels() ([]Channel, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/channels.list", nil)
	response := new(channelResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

func (r *Rocket) GetJoinedChannels() ([]Channel, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/channels.list.joined", nil)
	response := new(channelResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

func (r *Rocket) JoinChannel(channel *Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s" }`, channel.Id)
	request, _ := http.NewRequest("POST", r.getUrl() + "/api/v1/channels.join", bytes.NewBufferString(body))
	return r.doRequest(request, new(statusResponse))
}

func (r *Rocket) LeaveChannel(channel *Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.Id)
	request, _ := http.NewRequest("POST", r.getUrl() + "/api/v1/channels.leave", bytes.NewBufferString(body))
	return r.doRequest(request, new(statusResponse))
}

