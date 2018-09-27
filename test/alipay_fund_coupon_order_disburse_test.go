package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundCouponOrderDisburse(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundCouponOrderDisburseRequest{
		OutOrderNo:       "20180901000000AFCOPP00001",
		DeductOutOrderNo: "",
		OutRequestNo:     "20180901000000AFCOPP00001",
		OrderTitle:       "发送红包",
		Amount:           "9.99",
		PayeeLogonId:     "dlnhdb4422@sandbox.com",
	}
	client.SendRequest(request.AlipayFundCouponOrderDisburseMethod, data)
}
