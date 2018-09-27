package request

type AlipayTradeQueryRequest struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

var (
	AlipayTradeQueryMethod = "alipay.trade.query"
)
