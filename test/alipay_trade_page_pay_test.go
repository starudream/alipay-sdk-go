package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradePagePay(t *testing.T) {
	client, config := NewClient()
	client.SetReturnUrl(config.ReturnUrl)
	client.SetNotifyUrl(config.NotifyUrl)
	data := &request.AlipayTradePagePayRequest{
		OutTradeNo:  "20180901000000FITP00001",
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		TotalAmount: "9.99",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		Body:        "内容测试测试测试测试测试测试测试测试测试",
	}
	response, err := client.RequestUrl(request.AlipayTradePagePayMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
