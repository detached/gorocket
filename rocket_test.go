package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	testProtocol = "http"
	testHost = "localhost"
	testPort = "3000"
	testUserName = "test"
	testUserEmail = "test@test.de"
	testPassword = "test"
	rocketClient *Rocket
)

func getDefaultClient(t *testing.T) *Rocket {

	if (rocketClient == nil) {
		rocketClient = getAuthenticatedClient(t, testUserName, testUserEmail, testPassword)
	}

	return rocketClient
}

func getAuthenticatedClient(t *testing.T, name, email, password string) *Rocket {
	client := Rocket{Protocol: testProtocol, Host: testHost, Port: testPort}
	credentials := UserCredentials{Name:name, Email:email, Password:password}

	regErr := client.RegisterUser(credentials)
	assert.Nil(t, regErr)

	loginErr := client.Login(credentials)
	assert.Nil(t, loginErr)

	return &client
}

func findMessage(messages []Message, user string, msg string) *Message {
	for _, m := range messages {
		if m.User.UserName == user && m.Text == msg {
			return &m
		}
	}

	return nil
}

func getRoom(rooms []Room, name string) *Room {
	for _, r := range rooms {
		if r.Name == name {
			return &r
		}
	}

	return nil
}