package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetVersions(t *testing.T) {
	rocket := Rocket{Protocol:testProtocol, Host:testHost, Port:testPort}

	versions, err := rocket.GetVersions()

	assert.Nil(t, err)
	assert.NotNil(t, versions)
	assert.NotEmpty(t, versions.Version)
}
