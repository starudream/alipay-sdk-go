package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundAuthOperationDetailQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOperationDetailQueryRequest{
		OutOrderNo:   "20180901000000AFAOF00001",
		OutRequestNo: "20180901000000AFAOF00001",
	}
	_, _ = client.SendRequest(request.AlipayFundAuthOperationDetailQueryMethod, data)
}
