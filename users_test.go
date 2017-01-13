package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_LoginLogout(t *testing.T) {
	client := getAuthenticatedClient(t, getRandomString(), getRandomEmail(), getRandomString())
	_, logoutErr := client.Logout()
	assert.Nil(t, logoutErr)

	channels, err := client.GetJoinedChannels()
	assert.Nil(t, channels)
	assert.NotNil(t, err)
}