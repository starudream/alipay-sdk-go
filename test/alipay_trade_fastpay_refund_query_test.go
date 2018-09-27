package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeFastpayRefundQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeFastpayRefundQueryRequest{
		OutTradeNo:   "20180901000000ATWP00001",
		OutRequestNo: "20180901000000ATWP00001",
	}
	client.SendRequest(request.AlipayTradeFastpayRefundQueryMethod, data)
}
