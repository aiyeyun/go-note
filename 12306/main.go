package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	url2 "net/url"
)

func main() {
	yanzhen()
	os.Exit(0)
	url := "https://kyfw.12306.cn/otn/login/init"
	cookies, err := getHeaderSetCookies(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cookies)
}

func getHeaderSetCookies(url string) ([]*http.Cookie, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil, err
	}

	defer resp.Body.Close()

	return resp.Cookies(), nil
}

func yanzhen()  {
	url := "https://kyfw.12306.cn/otn/login/init"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	fmt.Println(resp.Cookies())

	url = "https://kyfw.12306.cn/otn/login/loginAysnSuggest"
	url = "https://kyfw.12306.cn/passport/web/login"
	data := &url2.Values{}
	data.Set("username", "1234567@qq.com")
	data.Set("password", "1234567")
	data.Set("appid", `otn`)
	req, _ := http.NewRequest("POST", url, strings.NewReader(data.Encode()))

	req.Header.Set("Origin", "https://kyfw.12306.cn")
	req.Header.Set("Host", "kyfw.12306.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.99 Safari/537.36")
	req.Header.Set("Referer", "https://kyfw.12306.cn/otn/login/init")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for i := range resp.Cookies() {
		req.AddCookie(resp.Cookies()[i])
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}