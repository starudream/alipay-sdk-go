package test

import (
	"github.com/starudream/alipay-sdk-go/request"
	"testing"
)

func TestAlipayFundAuthOperationCancel(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOperationCancelRequest{
		OutOrderNo:   "20180901000000PA00001",
		OutRequestNo: "20180901000000PA00001",
		Remark:       "授权撤销",
	}
	client.SendRequest(request.AlipayFundAuthOperationCancelMethod, data)
}
