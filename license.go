package main

import (
	"io/ioutil"
	chttp "net/http"
	"strings"
)

var licenseOk = true

func licenseCheck() {
	resp, err := chttp.Get("https://hi-hi.cn/license/cow-logger.v1")
	if err != nil {

	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if strings.Index(string(buf), "ok") == -1 {
		licenseOk = false
	}
}
