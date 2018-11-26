package test

import (
	"testing"
	"time"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashCreate(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashCreateRequest{
		CouponName:   "XXX周年庆红包",
		PrizeType:    "random",
		TotalMoney:   "1000",
		TotalNum:     "100",
		PrizeMsg:     "XXX送您大红包",
		StartTime:    "NowTime",
		EndTime:      time.Now().AddDate(0, 1, 0).Format("2006-01-02 15:04:05"),
		MerchantLink: "http://127.0.0.1",
	}
	_, _ = client.SendRequest(request.AlipayMarketingCampaignCashCreateMethod, data)
}
