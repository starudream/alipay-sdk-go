package test

import (
	"github.com/starudream/alipay-sdk-go/request"
	"testing"
)

func TestAlipayFundAuthOperationDetailQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOperationDetailQueryRequest{
		OutOrderNo:   "20180901000000AFAOF00001",
		OutRequestNo: "20180901000000AFAOF00001",
	}
	client.SendRequest(request.AlipayFundAuthOperationDetailQueryMethod, data)
}
