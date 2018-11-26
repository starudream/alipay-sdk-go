package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashTrigger(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashTriggerRequest{
		CrowdNo:  "5PZx2Y5c55NlJV_FXl0V0_Wve9z3gpyqu-HzZaTrTFTMnSZ96O-zxUfKlHp5cxmx",
		LoginId:  "dlnhdb4422@sandbox.com",
		OutBizNo: "20180901000000AMCCT00001",
	}
	_, _ = client.SendRequest(request.AlipayMarketingCampaignCashTriggerMethod, data)
}
