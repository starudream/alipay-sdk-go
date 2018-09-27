package request

type AlipayTradeWapPayRequest struct {
	Body               string                        `json:"body,omitempty"`
	Subject            string                        `json:"subject"`
	OutTradeNo         string                        `json:"out_trade_no"`
	TimeoutExpress     string                        `json:"timeout_express,omitempty"`
	TimeExpire         string                        `json:"time_expire,omitempty"`
	TotalAmount        string                        `json:"total_amount"`
	AuthToken          string                        `json:"auth_token,omitempty"`
	ProductCode        string                        `json:"product_code"`
	GoodsType          string                        `json:"goods_type,omitempty"`
	PassbackParams     string                        `json:"passback_params,omitempty"`
	PromoParams        string                        `json:"promo_params,omitempty"`
	ExtendParams       AlipayTradeWapPayExtendParams `json:"extend_params,omitempty"`
	EnablePayChannels  string                        `json:"enable_pay_channels,omitempty"`
	DisablePayChannels string                        `json:"disable_pay_channels,omitempty"`
	StoreId            string                        `json:"store_id,omitempty"`
	QuitUrl            string                        `json:"quit_url,omitempty"`
	ExtUserInfo        string                        `json:"ext_user_info,omitempty"`
}

type AlipayTradeWapPayExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	NeedBuyerRealnamed   string `json:"needBuyerRealnamed,omitempty"`
	TransMemo            string `json:"TRANS_MEMO,omitempty"`
	HbFqNum              string `json:"hb_fq_num,omitempty"`
	HbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`
}

var (
	AlipayTradeWapPayMethod = "alipay.trade.wap.pay"
)
