package request

const AlipayFundCouponOrderAppPayMethod = "alipay.fund.coupon.order.app.pay"

type AlipayFundCouponOrderAppPayRequest struct {
	OutOrderNo   string `json:"out_order_no"`
	OutRequestNo string `json:"out_request_no"`
	OrderTitle   string `json:"order_title"`
	Amount       string `json:"amount"`
	PayTimeout   string `json:"pay_timeout,omitempty"`
	ExtraParam   string `json:"extra_param,omitempty"`
}
