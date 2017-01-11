package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetPublicChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_GetJoinedChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(channels, "general")
	err = rocket.JoinChannel(general)
	assert.Nil(t, err)

	channels, err = rocket.GetJoinedChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_LeaveChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")
	err = rocket.JoinChannel(general)
	assert.Nil(t, err)

	err = rocket.LeaveChannel(general)
	assert.Nil(t, err)
}