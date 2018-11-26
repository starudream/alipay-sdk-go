package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundCouponOrderPagePay(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundCouponOrderPagePayRequest{
		OutOrderNo:   "20180901000000AFCOPP00001",
		OutRequestNo: "20180901000000AFCOPP00001",
		OrderTitle:   "发送红包",
		Amount:       "9.99",
	}
	_, _ = client.SendRequest(request.AlipayFundCouponOrderPagePayMethod, data)
}
