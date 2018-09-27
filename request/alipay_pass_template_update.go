package request

const AlipayPassTemplateUpdateMethod = "alipay.pass.template.update"

type AlipayPassTemplateUpdateRequest struct {
	TplId      string `json:"tpl_id"`
	TplContent string `json:"tpl_content"`
}
