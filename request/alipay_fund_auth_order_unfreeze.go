package request

const AlipayFundAuthOrderUnfreezeMethod = "alipay.fund.auth.order.unfreeze"

type AlipayFundAuthOrderUnfreezeRequest struct {
	AuthNo       string `json:"auth_no"`
	OutRequestNo string `json:"out_request_no"`
	Amount       string `json:"amount"`
	Remark       string `json:"remark"`
	ExtraParam   string `json:"extra_param,omitempty"`
}
