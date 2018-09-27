package request

const AlipayUserInfoAuthMethod = "alipay.user.info.auth"

type AlipayUserInfoAuthRequest struct {
	Scopes []string `json:"scopes"`
	State  string   `json:"state"`
}
