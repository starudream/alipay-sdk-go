package request

type AlipayPassInstanceUpdateRequest struct {
	SerialNumber string `json:"serial_number"`
	ChannelId    string `json:"channel_id"`
	TplParams    string `json:"tpl_params,omitempty"`
	Status       string `json:"status,omitempty"`
	VerifyCode   string `json:"verify_code,omitempty"`
	VerifyType   string `json:"verify_type,omitempty"`
}

var (
	AlipayPassInstanceUpdateMethod = "alipay.pass.instance.update"
)
