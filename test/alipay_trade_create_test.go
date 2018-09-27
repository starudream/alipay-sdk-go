package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeCreate(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeCreateRequest{
		OutTradeNo:  "20180901000000ATC00001",
		TotalAmount: "9.99",
		Subject:     "标题测试测试测试测试测试测试测试测试测试",
		BuyerId:     "2088102169174184",
	}
	client.SendRequest(request.AlipayTradeCreateMethod, data)
}
