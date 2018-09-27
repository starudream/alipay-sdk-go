package test

import (
	"io/ioutil"

	"github.com/json-iterator/go"
	"github.com/starudream/alipay-sdk-go"
)

type Config struct {
	AppId           string `json:"app_id"`
	PartnerId       string `json:"partner_id"`
	SellerId        string `json:"seller_id"`
	ReturnUrl       string `json:"return_url"`
	NotifyUrl       string `json:"notify_url"`
	PublicKey       string `json:"public_key"`
	PrivateKey      string `json:"private_key"`
	AlipayPublicKey string `json:"alipay_public_key"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewClient() (*alipaysdk.ClientData, *Config) {
	config := ReadConfigFile()

	client := &alipaysdk.ClientData{}
	client.DefaultSimpleDevelopmentClient(
		config.AppId,
		config.PrivateKey,
		config.AlipayPublicKey,
	)
	return client, config
}

func ReadConfigFile() (*Config) {
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
