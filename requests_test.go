package requests

import (
	"fmt"
	"testing"
)

func TestEncodeURl(t *testing.T) {
	urls := "http://www.baidu.com/"
	params := make(map[string]string)
	params["params1"] = "Mozilla"
	params["params2"] = "Go"
	urls, _ = EncodeURL(urls, params)
	fmt.Println(urls)
}

func TestGet(t *testing.T) {
	url := "http://www.baidu.com/"
	//params, headers := nil, nil
	resp, _ := Get(url, nil, nil, 0)
	fmt.Println(resp)
}
