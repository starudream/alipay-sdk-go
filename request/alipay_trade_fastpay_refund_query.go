package request

const AlipayTradeFastpayRefundQueryMethod = "alipay.trade.fastpay.refund.query"

type AlipayTradeFastpayRefundQueryRequest struct {
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no"`
}
