package request

type AlipayPassTemplateUpdateRequest struct {
	TplId      string `json:"tpl_id"`
	TplContent string `json:"tpl_content"`
}

var (
	AlipayPassTemplateUpdateMethod = "alipay.pass.template.update"
)
