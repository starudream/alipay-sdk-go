package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeQueryRequest{
		OutTradeNo: "20180901000000ATWP00001",
	}
	client.SendRequest(request.AlipayTradeQueryMethod, data)
}
