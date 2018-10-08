package test

import (
	"testing"
)

func TestCheckAsyncNotification(t *testing.T) {
	client, config := NewClient()
	// 传递参数类似于 gmt_create=xxx&charset=UTF-8&...
	// value 可以为未经过 url.QueryUnescape 的值
	// 在 gin 中传递 c.Request.URL.RawQuery 即可
	t.Log(client.CheckAsyncNotification(config.AsyncNotification))
}
