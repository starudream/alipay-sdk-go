package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeFastpayRefundQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeFastpayRefundQueryRequest{
		OutTradeNo:   "20180901000000QWW00001",
		OutRequestNo: "20180901000000QWW00001",
	}
	response, err := client.SendRequest(request.AlipayTradeFastpayRefundQueryMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
