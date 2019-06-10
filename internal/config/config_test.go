package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Load(t *testing.T) {
	// arrange
	data := `
---
http:
    port: 1234

log:
    level: DEBUG
    file: api.log

sentry:
  enabled: true
  key: https://xxxxx:yyyyy@sentry.io/123456

services:
  artists: https://internal/artists
  feed: https://internal/feed
  subscriptions: https://internal/subscriptions
  users: https://internal/users
  auth: https://internal/auth
`
	expected := &AppConfig{
		HTTP: HTTPConfig{
			Port: 1234,
		},
		Log: LogConfig{
			Level: "DEBUG",
			File:  "api.log",
		},
		Sentry: Sentry{
			Enabled: true,
			Key:     "https://xxxxx:yyyyy@sentry.io/123456",
		},
		Services: Services{
			Artists:       "https://internal/artists",
			Feed:          "https://internal/feed",
			Subscriptions: "https://internal/subscriptions",
			Users:         "https://internal/users",
			Auth:          "https://internal/auth",
		},
	}

	// action
	err := Load([]byte(data))

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, Config)
}
