package request

const AlipayTradeCloseMethod = "alipay.trade.close"

type AlipayTradeCloseRequest struct {
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OperatorId string `json:"operator_id"`
}
