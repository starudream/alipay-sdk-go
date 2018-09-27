package request

const AlipayOpenAuthTokenAppQueryMethod = "alipay.open.auth.token.app.query"

type AlipayOpenAuthTokenAppQueryRequest struct {
	AppAuthToken string `json:"app_auth_token"`
}
