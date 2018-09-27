package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeRefund(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeRefundRequest{
		OutTradeNo:   "20180901000000QWW00001",
		RefundAmount: "9.99",
		RefundReason: "退款测试",
		GoodsDetail: []request.AlipayTradeRefundGoodsDetail{
			{
				GoodsId:   "test",
				GoodsName: "test",
				Quantity:  "1",
				Price:     "19.99",
			},
		},
	}
	client.SendRequest(request.AlipayTradeRefundMethod, data)
}
