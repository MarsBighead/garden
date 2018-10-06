package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	filename := "/Users/duanp/config.xlsx"
	parseData(filename)

}

//parseData parse configure data from Excel
func parseData(filename string) (int, error) {
	xlsxFile, err := xlsx.OpenFile(filename)
	log.Println(err)
	if err != nil {
		return 0, err
	}
	for _, sheet := range xlsxFile.Sheets {
		for y, row := range sheet.Rows {
			if y == 0 {
				header := make(map[int]string)
				for x, cell := range row.Cells {
					val := cell.String()
					header[x] = val
					fmt.Println(val)
				}
				//u.ExtractBpHeader(y, row)
			} else {
				for _, cell := range row.Cells {
					val := cell.String()
					if regexp.MustCompile(`^https`).MatchString(val) {
						fmt.Println(val)
						get(val)
						os.Exit(0)
					}
				}
				//u.ExtractBpData(y, row)
			}
		}
	}
	return 0, nil
}
func login() {
	cookie, _ := cookiejar.New(nil)
	// Create client
	client := &http.Client{Jar: cookie}
	// Create request
	req, err := http.NewRequest("GET", "https://www.wjx.cn/login.aspx", nil)
	// Fetch Request
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	//开始修改缓存jar里面的值
	var cookieList []*http.Cookie
	cookieList = append(cookieList, &http.Cookie{
		Name:    "Wenjuanxing",
		Domain:  ".wjx.cn",
		Path:    "/",
		Value:   "cookie  值xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Expires: time.Now().AddDate(1, 0, 0),
	})
	uri, _ := url.Parse("https://www.wjx.cn/login.aspx")
	cookie.SetCookies(uri, cookieList)

	fmt.Printf("Jar cookie : %v", cookie.Cookies(uri))
	// Fetch Request
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
	fmt.Printf("response Cookies :%v", resp.Cookies())
}

func get(uri string) {
	re, _ := regexp.Compile(`*\.jpg`)
	name := re.FindStringSubmatch(uri)[0]
	//fmt.Print(name)
	//通过http请求获取图片的流文件
	resp, _ := http.Get(uri)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name)
	io.Copy(out, bytes.NewReader(body))

}
