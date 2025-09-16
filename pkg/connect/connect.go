package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: time.Second * 2,
}

// Get 判断url是否能够请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect.Get failed", logx.Field("url", url), logx.Field("error", err.Error()))
		return false
	}
	defer resp.Body.Close()
	// 200 OK
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
