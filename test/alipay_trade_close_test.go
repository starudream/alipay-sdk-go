package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeClose(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeCloseRequest{
		OutTradeNo: "20180901000000ATWP00001",
	}
	client.SendRequest(request.AlipayTradeCloseMethod, data)
}
