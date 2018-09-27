package request

const AlipayTradeQueryMethod = "alipay.trade.query"

type AlipayTradeQueryRequest struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	OrgPid     string `json:"org_pid,omitempty"`
}
