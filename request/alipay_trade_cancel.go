package request

const AlipayTradeCancelMethod = "alipay.trade.cancel"

type AlipayTradeCancelRequest struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}
