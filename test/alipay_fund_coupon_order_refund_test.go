package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundCouponOrderRefund(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundCouponOrderRefundRequest{
		AuthNo:       "",
		OutRequestNo: "20180901000000AFCOPP00001",
		Amount:       "9.99",
		Remark:       "红包退款",
	}
	client.SendRequest(request.AlipayFundCouponOrderRefundMethod, data)
}
