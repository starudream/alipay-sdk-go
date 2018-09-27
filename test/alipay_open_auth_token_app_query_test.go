package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayOpenAuthTokenAppQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayOpenAuthTokenAppQueryRequest{
		AppAuthToken: "201808BBef00aa93df16493fb6d0388e715f8X96",
	}
	client.SendRequest(request.AlipayOpenAuthTokenAppQueryMethod, data)
}
