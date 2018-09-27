package test

import (
	"testing"
)

func TestAlipayOauthAppToAppAuthUrl(t *testing.T) {
	client, config := NewClient()
	response := client.DefaultDevelopmentAppToAppAuthUrl(
		config.AppId,
		"http://127.0.0.1",
	)
	t.Log(response)
}
