package request

const AlipayPassInstanceAddMethod = "alipay.pass.instance.add"

type AlipayPassInstanceAddRequest struct {
	TplId           string `json:"tpl_id"`
	TplParams       string `json:"tpl_params"`
	RecognitionType string `json:"recognition_type"`
	RecognitionInfo string `json:"recognition_info"`
}
