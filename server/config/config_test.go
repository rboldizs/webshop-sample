package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConf(t *testing.T) {

	confFile := "config.json"
	config := GetConfig()

	assert.Nil(t, config.InitConf(confFile))

	assert.NotEmpty(t, config.GetServerPort())
	assert.NotEmpty(t, config.GetServerToken())

	dummyText := "Hello kitty"
	os.Setenv("SERVER_TOKEN", dummyText)
	os.Setenv("SERVER_PORT", "1388")

	config.InitConf(confFile)

	assert.Equal(t, dummyText, config.GetServerToken())
	assert.Equal(t, uint16(1388), config.GetServerPort())

}
