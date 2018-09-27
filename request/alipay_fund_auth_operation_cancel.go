package request

const AlipayFundAuthOperationCancelMethod = "alipay.fund.auth.operation.cancel"

type AlipayFundAuthOperationCancelRequest struct {
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Remark       string `json:"remark"`
}
