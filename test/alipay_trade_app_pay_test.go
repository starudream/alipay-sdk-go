package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeAppPay(t *testing.T) {
	client, config := NewClient()
	client.SetNotifyUrl(config.NotifyUrl)
	data := request.AlipayTradeAppPayRequest{
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		OutTradeNo:  "20180901000000ATAP00001",
		TotalAmount: "9.99",
		ProductCode: "QUICK_MSECURITY_PAY",
	}
	_, _ = client.SendRequest(request.AlipayTradeAppPayMethod, data)
}
