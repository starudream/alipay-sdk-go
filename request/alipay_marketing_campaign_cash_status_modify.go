package request

const AlipayMarketingCampaignCashStatusModifyMethod = "alipay.marketing.campaign.cash.status.modify"

type AlipayMarketingCampaignCashStatusModifyRequest struct {
	CrowdNo    string `json:"crowd_no"`
	CampStatus string `json:"camp_status"`
}
