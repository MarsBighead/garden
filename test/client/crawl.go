package test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Crawl Send http request and get response data
func Crawl(uri string) (*http.Response, error) {
	host, _ := os.Hostname()
	userAgent := "Cilent" + "/3.0" + " (Appcoach, " + host + ")"
	fmt.Printf("userAgent:%s\n", userAgent)
	proxyURL, _ := url.Parse("http://proxy.appcoachs.net:7070")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: time.Duration(300 * time.Second),
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	//resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
