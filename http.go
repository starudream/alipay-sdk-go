package alipaysdk

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) (string, error) {
	return HttpRequest("GET", url, nil)
}

func HttpRequest(method, url string, value []byte) (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(value))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("response status error:" + response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
