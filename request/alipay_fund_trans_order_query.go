package request

const AlipayFundTransOrderQueryMethod = "alipay.fund.trans.order.query"

type AlipayFundTransOrderQueryRequest struct {
	OutBizNo string `json:"out_biz_no"`
	OrderId  string `json:"order_id"`
}
