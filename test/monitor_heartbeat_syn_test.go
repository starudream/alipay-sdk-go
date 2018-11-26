package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestMonitorHeartbeatSyn(t *testing.T) {
	client, _ := NewClient()
	data := request.MonitorHeartbeatSynRequest{}
	_, _ = client.SendRequest(request.MonitorHeartbeatSynMethod, data)
}
