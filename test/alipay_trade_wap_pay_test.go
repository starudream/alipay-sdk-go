package test

import (
	"fmt"
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeWapPay(t *testing.T) {
	client, _ := NewClient()
	client.SetReturnUrl("http://127.0.0.1")
	client.SetNotifyUrl("http://127.0.0.1")
	data := request.AlipayTradeWapPayRequest{
		Body:        "内容测试测试测试测试测试测试测试测试测试",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		OutTradeNo:  "20180901000000QWW00001",
		TotalAmount: "9.99",
		ProductCode: "QUICK_WAP_WAY",
		ExtendParams: request.AlipayTradeWapPayExtendParams{
			TransMemo: "测试",
		},
	}
	response, err := client.RequestUrl(request.AlipayTradeWapPayMethod, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
