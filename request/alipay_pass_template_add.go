package request

type AlipayPassTemplateAddRequest struct {
	UniqueId   string `json:"unique_id"`
	TplContent string `json:"tpl_content"`
}

var (
	AlipayPassTemplateAddMethod = "alipay.pass.template.add"
)
