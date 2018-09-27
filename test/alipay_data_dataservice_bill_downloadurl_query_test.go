package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayDataDataserviceBillDownloadurlQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayDataDataserviceBillDownloadurlQueryRequest{
		BillType: "trade",
		BillDate: "2018-07",
	}
	client.SendRequest(request.AlipayDataDataserviceBillDownloadurlQueryMethod, data)
}
