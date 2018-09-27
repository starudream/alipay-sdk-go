package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipaySystemOauthToken(t *testing.T) {
	client, _ := NewClient()
	client.SetGrantTypeAndCode(
		"authorization_code",
		"2b744f7f07ba45f5804fbc9d89c4XX18",
	)
	response, err := client.SendRequest(request.AlipaySystemOauthTokenMethod, nil)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
