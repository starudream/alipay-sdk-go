package test

import (
	"testing"
)

func TestAlipayOauthAppToAppAuthUrl(t *testing.T) {
	client, config := NewClient()
	response := client.DefaultDevelopmentAppToAppAuthUrl(
		config.AppId,
		config.ReturnUrl,
	)
	t.Log(response)
}
