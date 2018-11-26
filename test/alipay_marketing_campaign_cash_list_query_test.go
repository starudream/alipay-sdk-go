package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashListQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashListQueryRequest{
		CampStatus: "ALL",
		PageSize:   "50",
		PageIndex:  "1",
	}
	_, _ = client.SendRequest(request.AlipayMarketingCampaignCashListQueryMethod, data)
}
