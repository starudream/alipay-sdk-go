package request

const AlipayFundCouponOperationQueryMethod = "alipay.fund.coupon.operation.query"

type AlipayFundCouponOperationQueryRequest struct {
	AuthNo       string `json:"auth_no"`
	OutOrderNo   string `json:"out_order_no"`
	OperationId  string `json:"operation_id"`
	OutRequestNo string `json:"out_request_no"`
}
