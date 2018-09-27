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
		OutTradeNo:  "20180901000000QMP00001",
		TotalAmount: "9.99",
		ProductCode: "QUICK_MSECURITY_PAY",
	}
	client.SendRequest(request.AlipayTradeAppPayMethod, data)
}
