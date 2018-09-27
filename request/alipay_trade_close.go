package request

type AlipayTradeCloseRequest struct {
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OperatorId string `json:"operator_id"`
}

var (
	AlipayTradeCloseMethod = "alipay.trade.close"
)
