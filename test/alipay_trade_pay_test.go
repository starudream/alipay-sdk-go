package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradePay(t *testing.T) {
	client, config := NewClient()
	client.SetNotifyUrl(config.NotifyUrl)
	data := &request.AlipayTradePayRequest{
		OutTradeNo:  "20180901000000ATP00001",
		Scene:       "bar_code",
		AuthCode:    "",
		ProductCode: "FACE_TO_FACE_PAYMENT",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		TotalAmount: "9.99",
		Body:        "内容测试测试测试测试测试测试测试测试测试",
	}
	client.SendRequest(request.AlipayTradePayMethod, data)
}
