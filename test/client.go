package test

import (
	"io/ioutil"

	"github.com/json-iterator/go"
	"github.com/starudream/alipay-sdk-go"
)

type Config struct {
	AppId           string `json:"app_id"`
	PrivateKey      string `json:"private_key"`
	AlipayPublicKey string `json:"alipay_public_key"`
	PartnerId       string `json:"partner_id"`
}

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

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
