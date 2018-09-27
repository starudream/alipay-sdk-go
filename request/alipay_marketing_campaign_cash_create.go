package request

type AlipayMarketingCampaignCashCreateRequest struct {
	CouponName   string `json:"coupon_name"`
	PrizeType    string `json:"prize_type"`
	TotalMoney   string `json:"total_money"`
	TotalNum     string `json:"total_num"`
	PrizeMsg     string `json:"prize_msg"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	MerchantLink string `json:"merchant_link,omitempty"`
	SendFreqency string `json:"send_freqency,omitempty"`
}

var (
	AlipayMarketingCampaignCashCreateMethod = "alipay.marketing.campaign.cash.create"
)
