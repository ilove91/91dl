package utils

import (
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

func GetNewHttpClient(timeout int64) *http.Client {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	str := viper.GetString("proxy")
	proxy, _ := url.ParseRequestURI(str)
	if proxy != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	} else {
		client.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	}

	return client
}
