package request

type AlipayTradePrecreateRequest struct {
	OutTradeNo           string                           `json:"out_trade_no"`
	SellerId             string                           `json:"seller_id,omitempty"`
	TotalAmount          string                           `json:"total_amount"`
	DiscountableAmount   string                           `json:"discountable_amount,omitempty"`
	Subject              string                           `json:"subject"`
	GoodsDetail          []AlipayTradePrecreateGoodDetail `json:"goods_detail,omitempty"`
	Body                 string                           `json:"body,omitempty"`
	OperatorId           string                           `json:"operator_id,omitempty"`
	StoreId              string                           `json:"store_id,omitempty"`
	DisablePayChannels   string                           `json:"disable_pay_channels"`
	EnablePayChannels    string                           `json:"enable_pay_channels"`
	TerminalId           string                           `json:"terminal_id,omitempty"`
	ExtendParams         AlipayTradePrecreateExtendParams `json:"extend_params,omitempty"`
	TimeoutExpress       string                           `json:"timeout_express,omitempty"`
	SettleInfo           string                           `json:"settle_info,omitempty"`
	BusinessParams       string                           `json:"business_params,omitempty"`
	QrCodeTimeoutExpress string                           `json:"qr_code_timeout_express,omitempty"`
}

type AlipayTradePrecreateGoodDetail struct {
	GoodsId       string `json:"goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category,omitempty"`
	Body          string `json:"body,omitempty"`
	ShowUrl       string `json:"show_url,omitempty"`
}

type AlipayTradePrecreateExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`
	CardType             string `json:"card_type,omitempty"`
}

type AlipayTradePrecreateSettleInfo struct {
	SettleDetailInfos []AlipayTradePrecreateSettleDetailInfo `json:"settle_detail_infos"`
	MerchantType      string                                 `json:"merchant_type,omitempty"`
}

type AlipayTradePrecreateSettleDetailInfo struct {
	TransInType      string `json:"trans_in_type"`
	TransIn          string `json:"trans_in"`
	SummaryDimension string `json:"summary_dimension,omitempty"`
	Amount           string `json:"amount"`
}

var (
	AlipayTradePrecreateMethod = "alipay.trade.precreate"
)
