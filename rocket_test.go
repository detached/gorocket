package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

var (
	testProtocol = "http"
	testHost = "localhost"
	testPort = "3000"
	testUserName string
	testUserEmail string
	testPassword = "test"
	rocketClient *Rocket
)

func getRandomString() string {
	length := 6
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func getRandomEmail() string {
	return getRandomString() + "@localhost.com"
}

func getDefaultClient(t *testing.T) *Rocket {

	if rocketClient == nil {
		testUserEmail = getRandomEmail()
		testUserName = getRandomString()
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

func getChannel(channels []Channel, name string) *Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}