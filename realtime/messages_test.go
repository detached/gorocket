package realtime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/detached/gorocket/api"
	"fmt"
)

func TestClient_SubscribeToMessageStream(t *testing.T) {

	c := getLoggedInClient(t)

	channel := api.Channel{Id: "GENERAL"}
	textToSend := "RealtimeTest"

	messageChannel, err := c.SubscribeToMessageStream(&channel)

	assert.Nil(t, err, "Function returned error")
	assert.NotNil(t, messageChannel, "Function didn't returned channel")

	c.SendMessage(&channel, textToSend)
	receivedMessage := <-messageChannel

	assert.NotNil(t, receivedMessage.Id, "Id was not set")
	assert.Equal(t, channel.Id, receivedMessage.ChannelId,"Wrong channel id")
	assert.NotNil(t, receivedMessage.Timestamp, "Timestamp was not set")
	assert.NotNil(t, receivedMessage.User.Id, "UserId was not set")
	assert.NotNil(t, receivedMessage.User.UserName, "Username was not set")
}

func TestClient_SubscribeToMessageStream_UnknownChannel(t *testing.T) {

	c := getLoggedInClient(t)
	channel := api.Channel{Id: "unknown"}

	messageChannel, err := c.SubscribeToMessageStream(&channel)
	fmt.Println("Subscribe done")

	assert.NotNil(t, err, "Function didn't return error")
	assert.Nil(t, messageChannel, "Function returned channel, but shouldn't")
}