package request

const AlipayFundCouponOrderDisburseMethod = "alipay.fund.coupon.order.disburse"

type AlipayFundCouponOrderDisburseRequest struct {
	OutOrderNo       string `json:"out_order_no"`
	DeductAuthNo     string `json:"deduct_auth_no,omitempty"`
	DeductOutOrderNo string `json:"deduct_out_order_no,omitempty"`
	OutRequestNo     string `json:"out_request_no"`
	OrderTitle       string `json:"order_title"`
	Amount           string `json:"amount"`
	PayeeUserId      string `json:"payee_user_id,omitempty"`
	PayeeLogonId     string `json:"payee_logon_id,omitempty"`
	PayTimeout       string `json:"pay_timeout,omitempty"`
	ExtraParam       string `json:"extra_param,omitempty"`
}
