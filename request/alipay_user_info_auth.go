package request

type AlipayUserInfoAuthRequest struct {
	Scopes []string `json:"scopes"`
	State  string   `json:"state"`
}

var (
	AlipayUserInfoAuthMethod = "alipay.user.info.auth"
)
