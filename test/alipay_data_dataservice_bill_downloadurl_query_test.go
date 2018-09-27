package test

import (
	"fmt"
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayDataDataserviceBillDownloadurlQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayDataDataserviceBillDownloadurlQueryRequest{
		BillType: "trade",
		BillDate: "2018-07",
	}
	response, err := client.SendRequest(request.AlipayDataDataserviceBillDownloadurlQueryMethod, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
