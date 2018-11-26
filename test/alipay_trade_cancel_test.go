package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeCancel(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeCancelRequest{
		OutTradeNo: "20180901000000ATWP00001",
	}
	_, _ = client.SendRequest(request.AlipayTradeCancelMethod, data)
}
