package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidferlay/gorocket/api"
)

type channelsResponse struct {
	Success  bool          `json:"success"`
	Channels []api.Channel `json:"channels"`
}

type groupsResponse struct {
	Succes bool        `json:"success"`
	Groups []api.Group `json:"groups"`
}

type channelResponse struct {
	Success bool        `json:"success"`
	Channel api.Channel `json:"channel"`
}

type groupResponse struct {
	Success bool      `json:"success"`
	Group   api.Group `json:"group"`
}

type groupUsersReponse struct {
	Success bool       `json:"success"`
	Members []api.User `json:"members"`
}

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) GetPublicChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) GetJoinedChannels() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/channels.list.joined", nil)
	response := new(channelsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all groups that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/groups/list
func (c *Client) GetJoinedGroups() ([]api.Group, error) {
	request, _ := http.NewRequest("GET", c.getUrl()+"/api/v1/groups.list", nil)
	response := new(groupsResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Groups, nil
}

// Joins a channel. The id of the channel has to be not nil.
//
// This function is not supported by the current Client.Chat release version 0.48.2.
func (c *Client) JoinChannel(channel *api.Channel) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
	}{
		RoomID: channel.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/channels.join", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

// Creates a group.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/create
func (c *Client) CreateGroup(group *api.Group) error {
	body, err := json.Marshal(struct {
		Name     string   `json:"name"`
		Members  []string `json:"members"`
		ReadOnly bool     `json:"readOnly"`
	}{
		Name:     group.Name,
		Members:  group.UserNames,
		ReadOnly: group.ReadOnly,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.create", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

// Archives a channel. The roomId has to be set.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/archive
func (c *Client) ArchiveGroup(group *api.Group) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
	}{
		RoomID: group.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.archive", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

// Unarchives a group. The roomId has to be set.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/unarchive
func (c *Client) UnarchiveGroup(group *api.Group) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
	}{
		RoomID: group.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.unarchive", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *api.Channel) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
	}{
		RoomID: channel.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/channels.leave", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

func (c *Client) KickFromGroup(group *api.Group, user *api.User) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
		UserID string `json:"userId"`
	}{
		RoomID: group.Id,
		UserID: user.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.kick", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

func (c *Client) KickFromChannel(channel *api.Channel, user *api.User) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
		UserID string `json:"userId"`
	}{
		RoomID: channel.Id,
		UserID: user.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/channels.kick", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Client) GetChannelInfo(channel *api.Channel) (*api.Channel, error) {
	var url = fmt.Sprintf("%s/api/v1/channels.info?roomId=%s&roomName=%s", c.getUrl(), channel.Id, channel.Name)
	request, _ := http.NewRequest("GET", url, nil)
	response := new(channelResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

// Invites a user to a channel
//
// https://rocket.chat/docs/developer-guides/rest-api/groups/invite
func (c *Client) InviteUser(group *api.Group, user *api.User) error {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId"`
		UserID string `json:"userId"`
	}{
		RoomID: group.Id,
		UserID: user.Id,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.invite", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

func (c *Client) GetGroupInfo(group *api.Group) (*api.Group, error) {
	url := c.getUrl() + "/api/v1/groups.info?"
	if group.Id != "" {
		url += "roomId=" + group.Id
	} else {
		url += "roomName=" + group.Name
	}
	request, _ := http.NewRequest("GET", url, nil)
	response := new(groupResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Group, nil
}

func (c *Client) GetGroupMembers(group *api.Group) (*[]api.User, error) {
	url := c.getUrl() + "/api/v1/groups.members?"
	if group.Id != "" {
		url += "roomId=" + group.Id
	} else {
		url += "roomName=" + group.Name
	}
	request, _ := http.NewRequest("GET", url, nil)
	response := new(groupUsersReponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Members, nil
}

func (c *Client) DeleteGroup(group *api.Group) error {
	body, err := json.Marshal(struct {
		RoomID   string `json:"roomId,omitempty"`
		RoomName string `json:"roomName,omitempty"`
	}{
		RoomID:   group.Id,
		RoomName: group.Name,
	})
	if err != nil {
		return err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.delete", bytes.NewReader(body))
	return c.doRequest(request, new(statusResponse))
}

func (c *Client) RenameGroup(group *api.Group) (*api.Group, error) {
	body, err := json.Marshal(struct {
		RoomID string `json:"roomId,omitempty"`
		Name   string `json:"name,omitempty"`
	}{
		RoomID: group.Id,
		Name:   group.Name,
	})
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPost, c.getUrl()+"/api/v1/groups.rename", bytes.NewReader(body))
	response := new(groupResponse)
	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}
	return &response.Group, nil
}
