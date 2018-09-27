package test

import (
	"testing"
)

func TestAlipayOauthPublicAppAuthUrl(t *testing.T) {
	client, config := NewClient()
	response := client.DefaultDevelopmentPublicAppAuthUrl(
		config.AppId,
		"http://127.0.0.1",
		"12345678",
		"auth_user",
	)
	t.Log(response)
}
