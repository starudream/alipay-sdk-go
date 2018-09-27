package request

const AlipayTradeAppPayMethod = "alipay.trade.app.pay"

type AlipayTradeAppPayRequest struct {
	Body               string                        `json:"body,omitempty"`
	Subject            string                        `json:"subject"`
	OutTradeNo         string                        `json:"out_trade_no"`
	TimeoutExpress     string                        `json:"timeout_express,omitempty"`
	TotalAmount        string                        `json:"total_amount"`
	ProductCode        string                        `json:"product_code"`
	GoodsType          string                        `json:"goods_type,omitempty"`
	PassbackParams     string                        `json:"passback_params,omitempty"`
	PromoParams        string                        `json:"promo_params,omitempty"`
	ExtendParams       AlipayTradeAppPayExtendParams `json:"extend_params,omitempty"`
	EnablePayChannels  string                        `json:"enable_pay_channels,omitempty"`
	DisablePayChannels string                        `json:"disable_pay_channels,omitempty"`
	StoreId            string                        `json:"store_id,omitempty"`
	ExtUserInfo        AlipayTradeAppPayExtUserInfo  `json:"ext_user_info,omitempty"`
}

type AlipayTradeAppPayExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	NeedBuyerRealnamed   string `json:"needBuyerRealnamed,omitempty"`
	TransMemo            string `json:"TRANS_MEMO,omitempty"`
	HbFqNum              string `json:"hb_fq_num,omitempty"`
	HbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`
}

type AlipayTradeAppPayExtUserInfo struct {
	Name          string `json:"name,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	CertType      string `json:"cert_type,omitempty"`
	CertNo        string `json:"cert_no,omitempty"`
	MinAge        string `json:"min_age,omitempty"`
	FixBuyer      string `json:"fix_buyer,omitempty"`
	NeedCheckInfo string `json:"need_check_info,omitempty"`
}
