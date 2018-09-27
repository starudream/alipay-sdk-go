package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradePay(t *testing.T) {
	client, _ := NewClient()
	client.SetNotifyUrl("http://127.0.0.1")
	data := &request.AlipayTradePayRequest{
		OutTradeNo:  "20180901000000FCFP00001",
		Scene:       "bar_code",
		AuthCode:    "",
		ProductCode: "FACE_TO_FACE_PAYMENT",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		TotalAmount: "9.99",
		Body:        "内容测试测试测试测试测试测试测试测试测试",
	}
	response, err := client.SendRequest(request.AlipayTradePayMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
