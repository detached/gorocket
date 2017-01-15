package realtime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/detached/gorocket/api"
)

func TestClient_SubscribeToMessageStream(t *testing.T) {

	c := getLoggedInClient(t)

	channel := api.Channel{Id: "GENERAL"}
	textToSend := "RealtimeTest"

	messageChannel, err := c.SubscribeToMessageStream(&channel)

	assert.NotNil(t, messageChannel)
	assert.Nil(t, err)

	c.SendMessage(&channel, textToSend)
	receivedMessage := <-messageChannel

	assert.NotNil(t, receivedMessage)
	assert.NotNil(t, receivedMessage.Id)
	assert.Equal(t, receivedMessage.ChannelId, channel.Id)
	assert.NotNil(t, receivedMessage.Timestamp)
	assert.NotNil(t, receivedMessage.User.Id)
	assert.NotNil(t, receivedMessage.User.UserName)
}
