package request

type AlipayOpenAuthTokenAppQueryRequest struct {
	AppAuthToken string `json:"app_auth_token"`
}

var (
	AlipayOpenAuthTokenAppQueryMethod = "alipay.open.auth.token.app.query"
)
