package request

const AlipayFundAuthOrderFreezeMethod = "alipay.fund.auth.order.freeze"

type AlipayFundAuthOrderFreezeRequest struct {
	AuthCode     string `json:"auth_code"`
	AuthCodeType string `json:"auth_code_type"`
	OutOrderNo   string `json:"out_order_no"`
	OutRequestNo string `json:"out_request_no"`
	OrderTitle   string `json:"order_title"`
	Amount       string `json:"amount"`
	PayeeLogonId string `json:"payee_logon_id,omitempty"`
	PayeeUserId  string `json:"payee_user_id,omitempty"`
	PayTimeout   string `json:"pay_timeout,omitempty"`
	ExtraParam   string `json:"extra_param,omitempty"`
	ProductCode  string `json:"product_code,omitempty"`
}
