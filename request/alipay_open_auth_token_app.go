package request

type AlipayOpenAuthTokenAppRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

var (
	AlipayOpenAuthTokenAppMethod = "alipay.open.auth.token.app"
)
