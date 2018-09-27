package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeCancel(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeCancelRequest{
		OutTradeNo: "20180901000000QWW00001",
	}
	response, err := client.SendRequest(request.AlipayTradeCancelMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
