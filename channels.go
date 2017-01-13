package gorocket

import (
	"fmt"
	"net/http"
	"bytes"
)

type channelsResponse struct {
	Success  bool `json:"success"`
	Channels []Channel `json:"channels"`
}

type channelResponse struct {
	Success bool `json:"success"`
	Channel Channel `json:"channel"`
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

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (r *Rocket) GetPublicChannels() ([]Channel, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/channels.list", nil)
	response := new(channelsResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (r *Rocket) GetJoinedChannels() ([]Channel, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/v1/channels.list.joined", nil)
	response := new(channelsResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Joins a channel. The id of the channel has to be not nil.
//
// This function is not supported by the current Rocket.Chat release version 0.48.2.
func (r *Rocket) JoinChannel(channel *Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s" }`, channel.Id)
	request, _ := http.NewRequest("POST", r.getUrl() + "/api/v1/channels.join", bytes.NewBufferString(body))
	return r.doRequest(request, new(statusResponse))
}

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (r *Rocket) LeaveChannel(channel *Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.Id)
	request, _ := http.NewRequest("POST", r.getUrl() + "/api/v1/channels.leave", bytes.NewBufferString(body))
	return r.doRequest(request, new(statusResponse))
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (r *Rocket) GetChannelInfo(channel *Channel) (*Channel, error) {
	var url = fmt.Sprintf("%s/api/v1/channels.info?roomId=%s", r.getUrl(), channel.Id)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(channelResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}