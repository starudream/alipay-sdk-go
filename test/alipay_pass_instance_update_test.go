package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayPassInstanceUpdate(t *testing.T) {
	client, config := NewClient()
	data := request.AlipayPassInstanceUpdateRequest{
		SerialNumber: "20180829090000",
		ChannelId:    config.AppId,
		TplParams:    `{"validStartDate":"2018-08-29 00:00:00","validEndDate":"2019-08-29 00:00:00","logoText":"测试优惠券00001","useStateDesc":"可用","discountInfo":"测试测试","channelID":"` + config.AppId + `","serialNumber":"20180829090000"}`,
		Status:       "CLOSED",
	}
	client.SendRequest(request.AlipayPassInstanceUpdateMethod, data)
}
