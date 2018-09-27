package request

const AlipayFundCouponOrderRefundMethod = "alipay.fund.coupon.order.refund"

type AlipayFundCouponOrderRefundRequest struct {
	AuthNo       string `json:"auth_no"`
	OutRequestNo string `json:"out_request_no"`
	Amount       string `json:"amount"`
	Remark       string `json:"remark"`
}
