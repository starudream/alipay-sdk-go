package test

import (
	"io/ioutil"

	"github.com/json-iterator/go"
	"github.com/starudream/alipay-sdk-go"
)

type Config struct {
	AppId             string `json:"app_id"`
	PartnerId         string `json:"partner_id"`
	SellerId          string `json:"seller_id"`
	ReturnUrl         string `json:"return_url"`
	NotifyUrl         string `json:"notify_url"`
	SignType          string `json:"sign_type"`
	PublicKey         string `json:"public_key"`
	PrivateKey        string `json:"private_key"`
	AlipayPublicKey   string `json:"alipay_public_key"`
	AsyncNotification string `json:"async_notification"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewClient() (*alipaysdk.ClientData, *Config) {
	alipaysdk.SetDebug(true)

	config := ReadConfigFile()

	client := &alipaysdk.ClientData{}
	client.DefaultProductionClient(
		config.AppId,
		"JSON",
		"UTF-8",
		config.SignType,
		config.PrivateKey,
		config.AlipayPublicKey,
	)
	return client, config
}

func ReadConfigFile() *Config {
	var data *Config

	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	return data
}
