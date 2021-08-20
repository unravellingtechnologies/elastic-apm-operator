package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLogLevelValid(t *testing.T) {
	level := parseLogLevel("DEBUG")
	assert.Equal(t, level, log.DebugLevel, "expected to find a log level of debug")

	level = parseLogLevel("error")
	assert.Equal(t, level, log.ErrorLevel, "expected to find a log level of error")
}

func TestParseLogLevelInvalid(t *testing.T) {
	level := parseLogLevel("war")
	assert.Equal(t, level, log.InfoLevel, "expected to find a log level of info")
}
