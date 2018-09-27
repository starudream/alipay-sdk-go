package test

import (
	"testing"
)

func TestAlipayOauthPublicAppAuthUrl(t *testing.T) {
	client, config := NewClient()
	response := client.DefaultDevelopmentPublicAppAuthUrl(
		config.AppId,
		config.ReturnUrl,
		"12345678",
		"auth_user",
	)
	t.Log(response)
}
