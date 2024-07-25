package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func VerifyHttp(pr string) bool {
	proxyUrl, proxyErr := url.Parse("http://" + pr)
	if proxyErr != nil {
		return false
	}
	tr := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	tr.Proxy = http.ProxyURL(proxyUrl)
	client := http.Client{Timeout: 10 * time.Second, Transport: &tr}
	request, err := http.NewRequest("GET", "http://baidu.com", nil)
	//处理返回结果
	res, err := client.Do(request)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	dataBytes, _ := io.ReadAll(res.Body)
	result := string(dataBytes)
	if strings.Contains(result, "0;url=http://www.baidu.com") {
		return true
	}
	return false
}
func VerifyHttps(pr string) bool {
	destConn, err := net.DialTimeout("tcp", pr, 10*time.Second)
	if err != nil {
		return false
	}
	defer destConn.Close()
	destConn.Write([]byte("CONNECT www.baidu.com:443 HTTP/1.1"))
	bytes := make([]byte, 1024)
	destConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	read, err := destConn.Read(bytes)
	fmt.Printf("%s\n", string(bytes[:read]))
	if strings.Contains(string(bytes[:read]), "200 Connection established") {
		return true
	}
	return false
}
