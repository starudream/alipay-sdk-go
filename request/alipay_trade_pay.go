package request

const AlipayTradePayMethod = "alipay.trade.pay"

type AlipayTradePayRequest struct {
	OutTradeNo         string                     `json:"out_trade_no"`
	Scene              string                     `json:"scene"`
	AuthCode           string                     `json:"auth_code"`
	ProductCode        string                     `json:"product_code,omitempty"`
	Subject            string                     `json:"subject"`
	BuyerId            string                     `json:"buyer_id,omitempty"`
	SellerId           string                     `json:"seller_id,omitempty"`
	TotalAmount        string                     `json:"total_amount,omitempty"`
	TransCurrency      string                     `json:"trans_currency,omitempty"`
	SettleCurrency     string                     `json:"settle_currency,omitempty"`
	DiscountableAmount string                     `json:"discountable_amount,omitempty"`
	Body               string                     `json:"body,omitempty"`
	GoodsDetail        []AlipayTradePayGoodDetail `json:"goods_detail,omitempty"`
	OperatorId         string                     `json:"operator_id,omitempty"`
	StoreId            string                     `json:"store_id,omitempty"`
	TerminalId         string                     `json:"terminal_id,omitempty"`
	ExtendParams       AlipayTradePayExtendParams `json:"extend_params,omitempty"`
	TimeoutExpress     string                     `json:"timeout_express,omitempty"`
	AuthConfirmMode    string                     `json:"auth_confirm_mode,omitempty"`
	TerminalParams     string                     `json:"terminal_params,omitempty"`
}

type AlipayTradePayGoodDetail struct {
	GoodsId       string `json:"goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category,omitempty"`
	Body          string `json:"body,omitempty"`
	ShowUrl       string `json:"show_url,omitempty"`
}

type AlipayTradePayExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`
	CardType             string `json:"card_type,omitempty"`
}
