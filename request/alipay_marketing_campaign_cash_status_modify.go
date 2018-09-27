package request

type AlipayMarketingCampaignCashStatusModifyRequest struct {
	CrowdNo    string `json:"crowd_no"`
	CampStatus string `json:"camp_status"`
}

var (
	AlipayMarketingCampaignCashStatusModifyMethod = "alipay.marketing.campaign.cash.status.modify"
)
