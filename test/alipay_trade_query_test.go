package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayTradeQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeQueryRequest{
		OutTradeNo: "20180901000000QWW00001",
	}
	response, err := client.SendRequest(request.AlipayTradeQueryMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
