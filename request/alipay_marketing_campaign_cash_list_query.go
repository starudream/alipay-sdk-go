package request

const AlipayMarketingCampaignCashListQueryMethod = "alipay.marketing.campaign.cash.list.query"

type AlipayMarketingCampaignCashListQueryRequest struct {
	CampStatus string `json:"camp_status,omitempty"`
	PageSize   string `json:"page_size"`
	PageIndex  string `json:"page_index"`
}
