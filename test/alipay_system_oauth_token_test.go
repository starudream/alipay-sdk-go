package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipaySystemOauthToken(t *testing.T) {
	client, _ := NewClient()
	_ = client.SetGrantTypeAndCode(
		"authorization_code",
		"ef929d7c82854897859d872547d6TX35",
	)
	_, _ = client.SendRequest(request.AlipaySystemOauthTokenMethod, nil)
}
