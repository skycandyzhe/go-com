package netcom

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

//添加url内容下载函数封装
func Httpsget(url string) (content string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	content = string(body)
	return content
}
