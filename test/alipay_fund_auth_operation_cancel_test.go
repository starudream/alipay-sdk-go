package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundAuthOperationCancel(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOperationCancelRequest{
		OutOrderNo:   "20180901000000AFAOF00001",
		OutRequestNo: "20180901000000AFAOF00001",
		Remark:       "授权撤销",
	}
	_, _ = client.SendRequest(request.AlipayFundAuthOperationCancelMethod, data)
}
