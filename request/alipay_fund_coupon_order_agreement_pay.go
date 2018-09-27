package request

const AlipayFundCouponOrderAgreementPayMethod = "alipay.fund.coupon.order.agreement.pay"

type AlipayFundCouponOrderAgreementPayRequest struct {
	OutOrderNo   string `json:"out_order_no"`
	OutRequestNo string `json:"out_request_no"`
	OrderTitle   string `json:"order_title"`
	Amount       string `json:"amount"`
	PayerUserId  string `json:"payer_user_id"`
	PayTimeout   string `json:"pay_timeout,omitempty"`
	ExtraParam   string `json:"extra_param,omitempty"`
}
