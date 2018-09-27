package request

const MonitorHeartbeatSynMethod = "monitor.heartbeat.syn"

type MonitorHeartbeatSynRequest struct {
	Product              string         `json:"product"`
	Type                 string         `json:"type"`
	EquipmentId          string         `json:"equipment_id"`
	Time                 string         `json:"time"`
	StoreId              string         `json:"store_id"`
	NetworkType          string         `json:"network_type"`
	SysServiceProviderId string         `json:"sys_service_provider_id,omitempty"`
	Mac                  string         `json:"mac,omitempty"`
	TradeInfo            []MHSTradeInfo `json:"trade_info"`
	ExceptionInfo        string         `json:"exception_info,omitempty"`
	ExtendInfo           string         `json:"extend_info,omitempty"`
}

type MHSTradeInfo struct {
	OTN  string `json:"OTN"`
	TC   string `json:"TC"`
	STAT string `json:"STAT"`
}
