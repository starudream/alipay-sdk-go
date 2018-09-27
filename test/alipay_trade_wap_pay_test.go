package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeWapPay(t *testing.T) {
	client, config := NewClient()
	client.SetReturnUrl(config.ReturnUrl)
	client.SetNotifyUrl(config.NotifyUrl)
	data := request.AlipayTradeWapPayRequest{
		Body:        "内容测试测试测试测试测试测试测试测试测试",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		OutTradeNo:  "20180901000000QWW00001",
		TotalAmount: "9.99",
		ProductCode: "QUICK_WAP_WAY",
		ExtendParams: &request.ATWPExtendParams{
			TransMemo: "测试",
		},
	}
	client.RequestUrl(request.AlipayTradeWapPayMethod, data)
}
