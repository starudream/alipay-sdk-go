package request

type AlipayFundTransOrderQueryRequest struct {
	OutBizNo string `json:"out_biz_no"`
	OrderId  string `json:"order_id"`
}

var (
	AlipayFundTransOrderQueryMethod = "alipay.fund.trans.order.query"
)
