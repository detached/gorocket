package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetPublicRooms(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicRooms()
	assert.Nil(t, err)

	assert.Len(t, rooms, 1)
	assert.Equal(t, "general", rooms[0].Name)
}

func TestRocket_GetJoinedRooms(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicRooms()
	assert.Nil(t, err)

	general := getRoom(rooms, "general")
	err = rocket.JoinRoom(general)
	assert.Nil(t, err)

	rooms, err = rocket.GetJoinedRooms()
	assert.Nil(t, err)

	assert.Len(t, rooms, 1)
	assert.Equal(t, "general", rooms[0].Name)
}

func TestRocket_LeaveRoom(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicRooms()
	assert.Nil(t, err)

	general := getRoom(rooms, "general")
	err = rocket.JoinRoom(general)
	assert.Nil(t, err)

	err = rocket.LeaveRoom(general)
	assert.Nil(t, err)
}