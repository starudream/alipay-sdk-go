package test

import (
	"fmt"
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashDetailQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashDetailQueryRequest{
		CrowdNo: "5PZx2Y5c55NlJV_FXl0V0_Wve9z3gpyqu-HzZaTrTFTMnSZ96O-zxUfKlHp5cxmx",
	}
	response, err := client.SendRequest(request.AlipayMarketingCampaignCashDetailQueryMethod, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
