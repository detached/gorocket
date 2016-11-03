package gorocket

import (
	"fmt"
	"net/http"
)

type roomsResponse struct {
	statusResponse
	Rooms []Room `json:"rooms"`
}

type channelResponse struct {
	Success bool `json:"success"`
	Channel Room `json:"channel"`
}

type Room struct {
	Id           string `json:"_id"`
	Name         string `json:"name"`
	Topic        string `json:"topic"`
	MessageCount int `json:"msgs"`
	UserNames    []string `json:"usernames"`

	User         User `json:"u"`

	ReadOnly     bool `json:"ro"`
	Timestamp    string `json:"ts"`
	T            string `json:"t"`
	UpdatedAt    string `json:"_updatedAt"`
	LastMessage  string `json:"lm"`
}

func (r *Rocket) GetPublicRooms() ([]Room, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/publicRooms", nil)
	response := new(roomsResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Rooms, nil
}

func (r *Rocket) GetJoinedRooms() ([]Room, error) {
	request, _ := http.NewRequest("GET", r.getUrl() + "/api/joinedRooms", nil)
	response := new(roomsResponse)

	if err := r.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Rooms, nil
}

func (r *Rocket) JoinRoom(room *Room) error {
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/rooms/%s/join", r.getUrl(), room.Id), nil)
	return r.doRequest(request, new(statusResponse))
}

func (r *Rocket) LeaveRoom(room *Room) error {
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/rooms/%s/leave", r.getUrl(), room.Id), nil)
	return r.doRequest(request, new(statusResponse))
}

