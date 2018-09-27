package request

const AlipayTradeRefundMethod = "alipay.trade.refund"

type AlipayTradeRefundRequest struct {
	OutTradeNo              string                      `json:"out_trade_no,omitempty"`
	TradeNo                 string                      `json:"trade_no,omitempty"`
	RefundAmount            string                      `json:"refund_amount"`
	RefundCurrency          string                      `json:"refund_currency,omitempty"`
	RefundReason            string                      `json:"refund_reason,omitempty"`
	OutRequestNo            string                      `json:"out_request_no,omitempty"`
	OperatorId              string                      `json:"operator_id,omitempty"`
	StoreId                 string                      `json:"store_id,omitempty"`
	TerminalId              string                      `json:"terminal_id,omitempty"`
	GoodsDetail             []ATRGoodDetail             `json:"goods_detail,omitempty"`
	RefundRoyaltyParameters []ATRRefundRoyaltyParameter `json:"refund_royalty_parameters,omitempty"`
	OrgPid                  string                      `json:"org_pid,omitempty"`
}

type ATRGoodDetail struct {
	GoodsId       string `json:"goods_id"`
	AlipayGoodsId string `json:"alipay_goods_id,omitempty"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category,omitempty"`
	Body          string `json:"body,omitempty"`
	ShowUrl       string `json:"show_url,omitempty"`
}

type ATRRefundRoyaltyParameter struct {
	RoyaltyType      string `json:"royalty_type"`
	TransOut         string `json:"trans_out"`
	TransOutType     string `json:"trans_out_type"`
	TransInType      string `json:"trans_in_type"`
	TransIn          string `json:"trans_in"`
	Amount           string `json:"amount"`
	AmountPercentage string `json:"amount_percentage"`
	Desc             string `json:"desc"`
}
