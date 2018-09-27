package request

const AlipayOpenAuthTokenAppMethod = "alipay.open.auth.token.app"

type AlipayOpenAuthTokenAppRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
