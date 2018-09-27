package alipaysdk

import (
	"crypto"
	"errors"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

type Client interface {
	DefaultClient(serverUrl, appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData)
	DefaultProductionClient(appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData)
	DefaultDevelopmentClient(appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData)
	DefaultSimpleProductionClient(appId, privateKey, alipayPublicKey string) (*ClientData)
	DefaultSimpleDevelopmentClient(appId, privateKey, alipayPublicKey string) (*ClientData)

	DefaultPublicAppAuthUrl(authUrl, appId, redirectUri, state string, scopes string) (string)
	DefaultProductionPublicAppAuthUrl(appId, redirectUri, state string, scopes string) (string)
	DefaultDevelopmentPublicAppAuthUrl(appId, redirectUri, state string, scopes string) (string)

	DefaultAppToAppAuthUrl(authUrl, appId, redirectUri string) (string)
	DefaultProductionAppToAppAuthUrl(appId, redirectUri string) (string)
	DefaultDevelopmentAppToAppAuthUrl(appId, redirectUri string) (string)

	SetAppAuthToken(appAuthToken string)
	SetGrantTypeAndCode(grantType, code string) (error)
	SetAuthToken(authToken string)
	SetReturnUrl(returnUrl string)
	SetNotifyUrl(notifyUrl string)

	RequestUrl(method string, bizContent interface{}) (string, error)
	SendRequest(method string, bizContent interface{}) (string, error)
}

type ClientData struct {
	ServerUrl       string      `json:"server_url"`        // 支付宝网关
	PublicKey       string      `json:"public_key"`        // 应用公钥
	PrivateKey      string      `json:"private_key"`       // 应用私钥
	AlipayPublicKey string      `json:"alipay_public_key"` // 支付宝公钥
	RequestData     RequestData `json:"request_data"`      // 公共请求参数
}

type AuthClientData struct {
	AppId       string   `json:"app_id"`       // 应用 id
	AuthUrl     string   `json:"auth_url"`     // 授权登录地址
	RedirectUri string   `json:"redirect_uri"` // 授权回调地址
	State       string   `json:"state"`        // 随机字符串，回调时返回给服务器
	Scopes      []string `json:"scopes"`       // 接口授权值
}

type RequestData struct {
	// 公共参数
	AppId     string `json:"app_id"`         // 应用 id
	Method    string `json:"method"`         // 接口名称
	Format    string `json:"format"`         // 格式
	Charset   string `json:"charset"`        // 请求使用的编码
	SignType  string `json:"sign_type"`      // 生成签名的算法
	Sign      string `json:"sign,omitempty"` // 签名
	Timestamp string `json:"timestamp"`      // 时间
	Version   string `json:"version"`        // 版本

	// 应用授权
	AppAuthToken string `json:"app_auth_token,omitempty"`

	// 请求参数集合
	BizContent string `json:"biz_content,omitempty"`

	// 其他参数，按需使用
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AuthToken    string `json:"auth_token,omitempty"`
	ReturnUrl    string `json:"return_url,omitempty"` // 回调地址
	NotifyUrl    string `json:"notify_url,omitempty"` // 异步通知地址
}

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (client *ClientData) DefaultClient(serverUrl, appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData) {
	client.ServerUrl = serverUrl
	client.PrivateKey = privateKey
	client.AlipayPublicKey = alipayPublicKey
	client.RequestData.AppId = appId
	client.RequestData.Format = format
	client.RequestData.Charset = charset
	client.RequestData.SignType = signType
	client.RequestData.Version = "1.0"
	return client
}

func (client *ClientData) DefaultProductionClient(appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData) {
	return client.DefaultClient(ServerUrlProduction, appId, format, charset, signType, privateKey, alipayPublicKey)
}

func (client *ClientData) DefaultDevelopmentClient(appId, format, charset, signType, privateKey, alipayPublicKey string) (*ClientData) {
	return client.DefaultClient(ServerUrlDevelopment, appId, format, charset, signType, privateKey, alipayPublicKey)
}

func (client *ClientData) DefaultSimpleProductionClient(appId, privateKey, alipayPublicKey string) (*ClientData) {
	return client.DefaultClient(ServerUrlProduction, appId, "JSON", "UTF-8", "RSA2", privateKey, alipayPublicKey)
}

func (client *ClientData) DefaultSimpleDevelopmentClient(appId, privateKey, alipayPublicKey string) (*ClientData) {
	return client.DefaultClient(ServerUrlDevelopment, appId, "JSON", "UTF-8", "RSA2", privateKey, alipayPublicKey)
}

func (client *ClientData) DefaultPublicAppAuthUrl(authUrl, appId, redirectUri, state string, scopes string) (string) {
	return authUrl +
		"?app_id=" + url.QueryEscape(appId) +
		"&scope=" + url.QueryEscape(scopes) +
		"&redirect_uri=" + url.QueryEscape(redirectUri) +
		"&state=" + url.QueryEscape(state)
}

func (client *ClientData) DefaultProductionPublicAppAuthUrl(appId, redirectUri, state string, scopes string) (string) {
	return client.DefaultPublicAppAuthUrl(PublicAppAuthUrlProduction, appId, redirectUri, state, scopes)
}

func (client *ClientData) DefaultDevelopmentPublicAppAuthUrl(appId, redirectUri, state string, scopes string) (string) {
	return client.DefaultPublicAppAuthUrl(PublicAppAuthUrlDevelopment, appId, redirectUri, state, scopes)
}

func (client *ClientData) DefaultAppToAppAuthUrl(authUrl, appId, redirectUri string) (string) {
	return authUrl +
		"?app_id=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectUri)
}

func (client *ClientData) DefaultProductionAppToAppAuthUrl(appId, redirectUri string) (string) {
	return client.DefaultAppToAppAuthUrl(AppToAppAuthUrlProduction, appId, redirectUri)
}

func (client *ClientData) DefaultDevelopmentAppToAppAuthUrl(appId, redirectUri string) (string) {
	return client.DefaultAppToAppAuthUrl(AppToAppAuthUrlDevelopment, appId, redirectUri)
}

func (client *ClientData) SetAppAuthToken(appAuthToken string) {
	client.RequestData.AppAuthToken = appAuthToken
}

func (client *ClientData) SetGrantTypeAndCode(grantType, code string) (error) {
	if grantType == "authorization_code" {
		client.RequestData.Code = code
	} else if grantType == "refresh_token" {
		client.RequestData.RefreshToken = code
	} else {
		return errors.New("unsupported grant_type")
	}
	client.RequestData.GrantType = grantType

	return nil
}

func (client *ClientData) SetAuthToken(authToken string) {
	client.RequestData.AuthToken = authToken
}

func (client *ClientData) SetReturnUrl(returnUrl string) {
	client.RequestData.ReturnUrl = returnUrl
}

func (client *ClientData) SetNotifyUrl(notifyUrl string) {
	client.RequestData.NotifyUrl = notifyUrl
}

func (client *ClientData) Validate() (error) {
	if client.RequestData.SignType != "RSA" && client.RequestData.SignType != "RSA2" {
		return errors.New("unsupported sign_type")
	}

	if client.RequestData.Charset != "UTF-8" {
		return errors.New("unsupported charset")
	}

	return nil
}

func (client *ClientData) ComposeParameterString() (string, string, error) {
	// 将结构体转换为 map
	bytes, err := json.Marshal(client.RequestData)
	if err != nil {
		return "", "", err
	}
	requestDataMap := make(map[string]string)
	json.Unmarshal(bytes, &requestDataMap)

	// 遍历 map 将 key 取出来并按照 ascii 排序
	var keys []string
	for key := range requestDataMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 拼接字符串
	signString := ""   // 待签名参数串
	encodeString := "" // 经过 url 编码的参数串
	for _, key := range keys {
		signString += key + "=" + requestDataMap[key] + "&"
		encodeString += key + "=" + url.QueryEscape(requestDataMap[key]) + "&"
	}
	// 去掉最后一个 "&"
	signString = signString[:len(signString)-1]
	encodeString = encodeString[:len(encodeString)-1]

	return signString, encodeString, nil
}

func (client *ClientData) RequestUrl(method string, bizContent interface{}) (string, error) {
	// 基本验证
	err := client.Validate()
	if err != nil {
		return "", err
	}

	// 设置请求方法名称
	client.RequestData.Method = method
	// 设置当前时间
	client.RequestData.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	// 设置请求参数集合
	if bizContent != nil {
		content, err := json.Marshal(bizContent)
		if err != nil {
			return "", err
		}
		client.RequestData.BizContent = string(content)
	}

	// 组合参数
	signString, encodeString, err := client.ComposeParameterString()
	if err != nil {
		return "", err
	}

	// 签名
	sign := ""
	if client.RequestData.SignType == "RSA" {
		sign, err = RsaSign(signString, client.PrivateKey, crypto.SHA1)
	} else if client.RequestData.SignType == "RSA2" {
		sign, err = RsaSign(signString, client.PrivateKey, crypto.SHA256)
	}
	if err != nil {
		return signString, err
	}
	client.RequestData.Sign = sign

	// 拼接 url 并添加 sign 字段
	return client.ServerUrl + "?" + encodeString + "&sign=" + url.QueryEscape(sign), nil
}

func (client *ClientData) SendRequest(method string, bizContent interface{}) (string, error) {
	// 获取请求地址
	requestUrl, err := client.RequestUrl(method, bizContent)
	if err != nil {
		return "", err
	}

	// 发送请求
	body, err := GetRequest(requestUrl)
	if err != nil {
		return requestUrl, err
	}

	// 首先通过解析 json 判断有必要字段
	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	contentInterface := response["error_response"]
	if contentInterface == nil {
		contentInterface = response[strings.Replace(method, ".", "_", -1)+"_response"]
	}
	signInterface := response["sign"]
	if contentInterface == nil || signInterface == nil {
		return body, errors.New("response must have `method_response` and `sign` field")
	}

	// 通过 index 来取出相应的字段
	// 由于 golang 的无序性，解密加密会产生不同的结果，以至于无法正确验签
	contentStartIndex := strings.Index(body, `_response":`) + 11
	contentEndIndex := strings.LastIndex(body, `"sign":`) - 1
	contentJsonString := body[contentStartIndex:contentEndIndex]
	signString := body[contentEndIndex+9 : len(body)-2]

	// 验证签名
	if client.RequestData.SignType == "RSA" {
		err = RsaCheck(contentJsonString, signString, client.AlipayPublicKey, crypto.SHA1)
		if err != nil {
			return contentJsonString + "\n" + signString, err
		}
	} else if client.RequestData.SignType == "RSA2" {
		err = RsaCheck(contentJsonString, signString, client.AlipayPublicKey, crypto.SHA256)
		if err != nil {
			return contentJsonString + "\n" + signString, err
		}
	}

	return contentJsonString, nil
}
