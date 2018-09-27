package request

const AlipayPassTemplateAddMethod = "alipay.pass.template.add"

type AlipayPassTemplateAddRequest struct {
	UniqueId   string `json:"unique_id"`
	TplContent string `json:"tpl_content"`
}
