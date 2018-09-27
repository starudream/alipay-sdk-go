package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayUserInfoAuth(t *testing.T) {
	client, config := NewClient()
	client.SetReturnUrl(config.ReturnUrl)
	data := &request.AlipayUserInfoAuthRequest{
		Scopes: []string{"auth_user"},
		State:  "12345678",
	}
	response, err := client.RequestUrl(request.AlipayUserInfoAuthMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
