package realtime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/detached/gorocket/common_testing"
	"github.com/detached/gorocket/api"
)

var (
	client *Client
)

func getLoggedInClient(t *testing.T) *Client {

	if client == nil {
		c, err := NewClient(common_testing.Host, common_testing.Port, true)
		assert.Nil(t, err)

		err = c.RegisterUser(&api.UserCredentials{
			Email: common_testing.GetRandomEmail(),
			Name: common_testing.GetRandomString(),
			Password: common_testing.GetRandomString()})
		assert.Nil(t, err)

		client = c
	}

	return client
}