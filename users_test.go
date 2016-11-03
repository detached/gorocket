package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_LoginLogout(t *testing.T) {
	client := getAuthenticatedClient(t, "loginLogout", "login@logout.de", "loginLogout")

	_, logoutErr := client.Logout()
	assert.Nil(t, logoutErr)
}