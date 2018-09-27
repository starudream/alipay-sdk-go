package request

const AlipayMarketingCampaignCashTriggerMethod = "alipay.marketing.campaign.cash.trigger"

type AlipayMarketingCampaignCashTriggerRequest struct {
	UserId     string `json:"user_id,omitempty"`
	CrowdNo    string `json:"crowd_no"`
	LoginId    string `json:"login_id,omitempty"`
	OrderPrice string `json:"order_price,omitempty"`
	OutBizNo   string `json:"out_biz_no,omitempty"`
}
