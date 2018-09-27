package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashStatusModify(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashStatusModifyRequest{
		CrowdNo:    "5PZx2Y5c55NlJV_FXl0V0_Wve9z3gpyqu-HzZaTrTFTMnSZ96O-zxUfKlHp5cxmx",
		CampStatus: "PAUSE",
	}
	response, err := client.SendRequest(request.AlipayMarketingCampaignCashStatusModifyMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
