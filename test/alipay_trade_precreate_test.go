package test

import (
	"fmt"
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradePrecreate(t *testing.T) {
	client, _ := NewClient()
	client.SetNotifyUrl("http://127.0.0.1")
	data := &request.AlipayTradePrecreateRequest{
		OutTradeNo:  "20180901000000P00001",
		TotalAmount: "9.99",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		Body:        "内容测试测试测试测试测试测试测试测试测试",
	}
	response, err := client.SendRequest(request.AlipayTradePrecreateMethod, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
