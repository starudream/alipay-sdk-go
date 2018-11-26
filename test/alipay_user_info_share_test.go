package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayUserInfoShare(t *testing.T) {
	client, _ := NewClient()
	client.SetAuthToken("authusrBf5c3b3abb20c41bd83b0c96fea39eF18")
	_, _ = client.SendRequest(request.AlipayUserInfoShareMethod, nil)
}
