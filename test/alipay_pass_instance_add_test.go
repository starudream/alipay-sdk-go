package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayPassInstanceAdd(t *testing.T) {
	client, config := NewClient()
	data := request.AlipayPassInstanceAddRequest{
		TplId:           "2018082909010397439888491",
		TplParams:       `{"validStartDate":"2018-08-29 00:00:00","validEndDate":"2019-08-29 00:00:00","logoText":"测试优惠券00001","useStateDesc":"可用","discountInfo":"测试测试","channelID":"` + config.AppId + `","serialNumber":"20180829090000"}`,
		RecognitionType: "1",
		RecognitionInfo: `{"partner_id":"` + config.PartnerId + `","out_trade_no":"20180828150000_1012216190_a1b2c3d4e5"}`,
	}
	_, _ = client.SendRequest(request.AlipayPassInstanceAddMethod, data)
}
