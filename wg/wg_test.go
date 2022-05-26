package wg

import (
	"testing"
	"wg/config"

	"github.com/stretchr/testify/assert"
)

func init() {
	config.C = &config.Config{
		Path: "../wg0.conf",
	}
}

func TestGetPortNumber(t *testing.T) {
	t.Logf("config: %+v", config.C)
	assert.Equal(t, "../wg0.conf", config.C.Path)

	num, err := GetPortNumber()
	assert.NoError(t, err, "")

	assert.Equal(t, uint(51825), num)
}

func TestIncrPortNumber(t *testing.T) {
	t.Logf("config: %+v", config.C)

	curPort, err := GetPortNumber()
	assert.NoError(t, err, "")

	newPort, err := IncrPortNumber()
	assert.NoError(t, err, "")

	assert.Equal(t, curPort+1, newPort)
}
