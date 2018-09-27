package request

const AlipayFundAuthOperationDetailQueryMethod = "alipay.fund.auth.operation.detail.query"

type AlipayFundAuthOperationDetailQueryRequest struct {
	AuthNo       string `json:"auth_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
}
