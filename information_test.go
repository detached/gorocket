package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestRocket_GetServerInfo(t *testing.T) {
	rocket := Rocket{Protocol:testProtocol, Host:testHost, Port:testPort}

	info, err := rocket.GetServerInfo()

	assert.Nil(t, err)
	assert.NotNil(t, info)

	assert.NotEmpty(t, info.Version)

	assert.NotEmpty(t, info.Build.Arch)
	assert.NotZero(t, info.Build.CpuCount)
	assert.NotEmpty(t, info.Build.Platform)
	assert.NotEmpty(t, info.Build.Date)
	assert.NotZero(t, info.Build.FreeMemory)
	assert.NotZero(t, info.Build.TotalMemory)
	assert.NotEmpty(t, info.Build.NodeVersion)
	assert.NotEmpty(t, info.Build.OsRelease)

	assert.NotEmpty(t, info.Travis.Branch)
	assert.NotEmpty(t, info.Travis.BuildNumber)
	assert.NotEmpty(t, info.Travis.Tag)

	assert.NotEmpty(t, info.Commit.Author)
	assert.NotEmpty(t, info.Commit.Branch)
	assert.NotEmpty(t, info.Commit.Date)
	assert.NotEmpty(t, info.Commit.Hash)
	assert.NotEmpty(t, info.Commit.Subject)
	assert.NotEmpty(t, info.Commit.Tag)

	assert.NotEmpty(t, info.ImageMagick.Version)
	assert.NotNil(t, info.GraphicsMagick)
}
