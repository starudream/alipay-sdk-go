package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestMonitorHeartbeatSyn(t *testing.T) {
	client, _ := NewClient()
	data := request.MonitorHeartbeatSynRequest{}
	response, err := client.SendRequest(request.MonitorHeartbeatSynMethod, data)
	if err != nil {
		panic(err)
	}
	t.Log(response)
}
