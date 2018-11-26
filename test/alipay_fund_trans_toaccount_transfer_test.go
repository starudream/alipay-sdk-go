package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundTransToaccountTransfer(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundTransToaccountTransferRequest{
		OutBizNo:      "20180901000000AFTTT00001",
		PayeeType:     "ALIPAY_LOGONID",
		PayeeAccount:  "dlnhdb4422@sandbox.com",
		Amount:        "9.99",
		PayerShowName: "测试转账",
		PayeeRealName: "沙箱环境",
		Remark:        "测试转账测试转账",
	}
	_, _ = client.SendRequest(request.AlipayFundTransToaccountTransferMethod, data)
}
