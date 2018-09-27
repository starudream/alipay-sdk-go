package request

type AlipayDataDataserviceBillDownloadurlQueryRequest struct {
	BillType string `json:"bill_type"`
	BillDate string `json:"bill_date"`
}

var (
	AlipayDataDataserviceBillDownloadurlQueryMethod = "alipay.data.dataservice.bill.downloadurl.query"
)
