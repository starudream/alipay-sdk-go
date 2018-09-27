package request

type AlipayFundTransToaccountTransferRequest struct {
	OutBizNo      string `json:"out_biz_no"`
	PayeeType     string `json:"payee_type"`
	PayeeAccount  string `json:"payee_account"`
	Amount        string `json:"amount"`
	PayerShowName string `json:"payer_show_name,omitempty"`
	PayeeRealName string `json:"payee_real_name,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

var (
	AlipayFundTransToaccountTransferMethod = "alipay.fund.trans.toaccount.transfer"
)
