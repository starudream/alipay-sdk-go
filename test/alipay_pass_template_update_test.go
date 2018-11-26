package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayPassTemplateUpdate(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayPassTemplateUpdateRequest{
		TplId:      "2018082909010397439888491",
		TplContent: `{"logo":"https://alipass.alipay.com//temps/free/logo.png","strip":"https://alipass.alipay.com//temps/free/strip.png","icon":"http://alipassprod.test.alipay.net/temps/free/icon.png","content":{"evoucherInfo":{"title":"风情优惠券","type":"coupon","product":"free","startDate":"$validStartDate$","endDate":"$validEndDate$","operation":[{"format":"qrcode","message":"45612346579465","messageEncoding":"UTF-8","altText":"45612346579465"},{"format":"barcode","message":"45612346579465","messageEncoding":"UTF-8","altText":"45612346579465"}],"einfo":{"logoText":"$logoText$","headFields":[{"key":"status","label":"状态","value":"$useStateDesc$","type":"text"}],"primaryFields":[{"key":"strip","label":"","value":"$discountInfo$","type":"text"}],"secondaryFields":[{"key":"validDate","label":"有效期至：","value":"$validEndDate$","type":"text"}],"auxiliaryFields":[],"backFields":[{"key":"description","label":"详情描述","value":"1.该优惠有效期：截止至2014年06月18日；\n2.凭此券可以享受以下优惠：\n享门市价优惠\n不与其他优惠同享。详询商家。","type":"text"},{"key":"shops","label":"可用门店","value":"","type":"text"},{"key":"disclaimer","label":"负责声明","value":"除特殊注明外，本优惠不能与其他优惠同时享受； 本优惠最终解释权归商家所有，如有疑问请与商家联系。 提示：为了使您得到更好的服务，请在进店时出示本券。","type":"text"}]},"remindInfo":{"offset":"2"}},"merchant":{"mname":"hodewu","mtel":"","minfo":""},"platform":{"channelID":"$channelID$","webServiceUrl":"http://alipassprod.test.alipay.net/builder/syncRecord.htm?tempId=2016062716334639647179025"},"style":{"backgroundColor":"RGB(255,126,0)"},"fileInfo":{"formatVersion":"2","canShare":true,"canBuy":false,"canPresent":true,"serialNumber":"$serialNumber$","supportTaxi":"true","taxiSchemaUrl":"alipays://platformapi/startapp?appId=20000130&sourceId=20000030&showTitleBar=YES&showToolBar=NO&showLoading=NO&safePayEnabled=YES&readTitle=YES&backBehavior=back&url=/www/index.html"},"appInfo":{"app":{"android_appid":"com.taobao.ecoupon","ios_appid":"taobaolife://","android_launch":"com.taobao.ecoupon","ios_launch":"taobaolife://","android_download":"http://download.taobaocdn.com/freedom/17988/andriod/Ecoupon_2.0.1_taobao_wap.apk","ios_download":"https://itunes.apple.com/cn/app/id583295537"},"label":"淘宝券券APP","message":"点击调起APP"},"source":"alipassprod","alipayVerify":[]}}`,
	}
	_, _ = client.SendRequest(request.AlipayPassTemplateUpdateMethod, data)
}
