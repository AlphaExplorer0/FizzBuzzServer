package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// We can have only one test for config due to the use of environment variables
func TestParseConfig(t *testing.T) {
	config := Config{}
	err := config.Init()
	assert.Nil(t, err)
	// default config
	assert.Equal(t, Config{
		BindIP:   "0.0.0.0",
		BindPort: 8080,
	}, config)

	os.Setenv("FIZZBUZZ_BINDIP", "127.0.0.1")
	os.Setenv("FIZZBUZZ_BINDPORT", "4567")

	err = config.Init()
	assert.Nil(t, err)
	assert.Equal(t, Config{
		BindIP:   "127.0.0.1",
		BindPort: 4567,
	}, config)
}
