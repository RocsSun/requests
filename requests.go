package requests

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

// EncodeURL 通过params 给url添加参数，并编码
func EncodeURL(urls string, params map[string]string) (string, error) {

	tmpURL, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	paramsValues := url.Values{}

	if params != nil {
		for k, v := range params {
			paramsValues.Set(k, v)
		}
		tmpURL.RawQuery = paramsValues.Encode()
	}

	return tmpURL.String(), nil
}

// Get http get method
func Get(urls string, params, headers map[string]string, timeout int64) (*http.Response, error) {
	if urls == "" {
		return nil, errors.New("URL不能为空")
	}
	urls, err := EncodeURL(urls, params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = time.Duration(timeout)
	}

	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	return response, err
}

// Post http post method
func Post(urls, types string, headers map[string]string) {

}
