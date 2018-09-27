package request

const AlipayTradeCreateMethod = "alipay.trade.create"

type AlipayTradeCreateRequest struct {
	OutTradeNo          string                  `json:"out_trade_no"`
	SellerId            string                  `json:"seller_id,omitempty"`
	TotalAmount         string                  `json:"total_amount"`
	DiscountableAmount  string                  `json:"discountable_amount,omitempty"`
	Subject             string                  `json:"subject"`
	Body                string                  `json:"body,omitempty"`
	BuyerId             string                  `json:"buyer_id,omitempty"`
	GoodsDetail         []ATCGoodDetail         `json:"goods_detail,omitempty"`
	OperatorId          string                  `json:"operator_id,omitempty"`
	StoreId             string                  `json:"store_id,omitempty"`
	TerminalId          string                  `json:"terminal_id,omitempty"`
	ExtendParams        *ATCExtendParams        `json:"extend_params,omitempty"`
	TimeoutExpress      string                  `json:"timeout_express,omitempty"`
	SettleInfo          *ATCSettleInfo          `json:"settle_info,omitempty"`
	BusinessParams      string                  `json:"business_params,omitempty"`
	ReceiverAddressInfo *ATCReceiverAddressInfo `json:"receiver_address_info,omitempty"`
	LogisticsDetail     *ATCLogisticsDetail     `json:"logistics_detail,omitempty"`
}

type ATCGoodDetail struct {
	GoodsId       string `json:"goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category,omitempty"`
	Body          string `json:"body,omitempty"`
	ShowUrl       string `json:"show_url,omitempty"`
}

type ATCExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`
	CardType             string `json:"card_type,omitempty"`
}

type ATCSettleInfo struct {
	SettleDetailInfos []ATCSettleDetailInfo `json:"settle_detail_infos"`
	MerchantType      string                `json:"merchant_type,omitempty"`
}

type ATCSettleDetailInfo struct {
	TransInType      string `json:"trans_in_type"`
	TransIn          string `json:"trans_in"`
	SummaryDimension string `json:"summary_dimension,omitempty"`
	SettleEntityId   string `json:"settle_entity_id,omitempty"`
	SettleEntityType string `json:"settle_entity_type,omitempty"`
	Amount           string `json:"amount"`
}

type ATCReceiverAddressInfo struct {
	Name         string `json:"name,omitempty"`
	Address      string `json:"address,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Zip          string `json:"zip,omitempty"`
	DivisionCode string `json:"division_code,omitempty"`
}

type ATCLogisticsDetail struct {
	LogisticsType string `json:"logistics_type,omitempty"`
}
