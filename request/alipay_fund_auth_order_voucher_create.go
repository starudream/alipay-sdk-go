package request

const AlipayFundAuthOrderVoucherCreateMethod = "alipay.fund.auth.order.voucher.create"

type AlipayFundAuthOrderVoucherCreateRequest struct {
	OutOrderNo        string `json:"out_order_no"`
	OutRequestNo      string `json:"out_request_no"`
	OrderTitle        string `json:"order_title"`
	Amount            string `json:"amount"`
	PayeeLogonId      string `json:"payee_logon_id,omitempty"`
	PayeeUserId       string `json:"payee_user_id,omitempty"`
	PayTimeout        string `json:"pay_timeout,omitempty"`
	ExtraParam        string `json:"extra_param,omitempty"`
	ProductCode       string `json:"product_code,omitempty"`
	TransCurrency     string `json:"trans_currency,omitempty"`
	SettleCurrency    string `json:"settle_currency,omitempty"`
	EnablePayChannels string `json:"enable_pay_channels,omitempty"`
}
