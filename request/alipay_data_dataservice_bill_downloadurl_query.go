package request

const AlipayDataDataserviceBillDownloadurlQueryMethod = "alipay.data.dataservice.bill.downloadurl.query"

type AlipayDataDataserviceBillDownloadurlQueryRequest struct {
	BillType string `json:"bill_type"`
	BillDate string `json:"bill_date"`
}
