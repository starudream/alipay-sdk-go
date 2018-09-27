package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundAuthOrderVoucherCreate(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOrderVoucherCreateRequest{
		OutOrderNo:   "20180901000000AFAOVC00001",
		OutRequestNo: "20180901000000AFAOVC00001",
		OrderTitle:   "预授权发码",
		Amount:       "9.99",
		PayeeLogonId: "dlnhdb4422@sandbox.com",
	}
	client.SendRequest(request.AlipayFundAuthOrderVoucherCreateMethod, data)
}
