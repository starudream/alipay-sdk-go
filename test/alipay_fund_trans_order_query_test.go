package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundTransOrderQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundTransOrderQueryRequest{
		OutBizNo: "20180901000000AFTTT00001",
	}
	_, _ = client.SendRequest(request.AlipayFundTransOrderQueryMethod, data)
}
