package test

import (
	"fmt"
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeRefund(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeRefundRequest{
		OutTradeNo:   "20180901000000QWW00001",
		RefundAmount: "9.99",
		RefundReason: "退款测试",
	}
	response, err := client.SendRequest(request.AlipayTradeRefundMethod, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
