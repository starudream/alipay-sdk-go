package test

import (
	"github.com/starudream/alipay-sdk-go/request"
	"testing"
)

func TestAlipayFundAuthOrderFreeze(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOrderFreezeRequest{
		AuthCode:     "281442445657444252",
		AuthCodeType: "bar_code",
		OutOrderNo:   "20180901000000PA00001",
		OutRequestNo: "20180901000000PA00001",
		OrderTitle:   "预授权冻结",
		Amount:       "9.99",
		PayeeLogonId: "dlnhdb4422@sandbox.com",
	}
	client.SendRequest(request.AlipayFundAuthOrderFreezeMethod, data)
}
