package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url" // exported package is phantomjs
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	phantomjs "github.com/urturn/go-phantomjs"
)

//全局变量
var CurCookies []*http.Cookie
var CurCookieJar *cookiejar.Jar //管理cookie

//初始化
func init() {
	CurCookies = nil
	//var err error;
	CurCookieJar, _ = cookiejar.New(nil)
}

//get url response html
func getUrlRespHtml(strUrl string, postDict map[string]string) string {
	fmt.Printf("in getUrlRespHtml, strUrl=%s\n", strUrl)
	fmt.Printf("postDict=%s\n", postDict)

	var respHtml string = ""

	httpClient := &http.Client{
		Jar: CurCookieJar,
	}

	var httpReq *http.Request
	if nil == postDict {
		fmt.Printf("is GET\n")
		httpReq, _ = http.NewRequest("GET", strUrl, nil)

	} else {
		fmt.Printf("is POST\n")
		postValues := url.Values{}
		for postKey, PostValue := range postDict {
			postValues.Set(postKey, PostValue)
		}
		fmt.Printf("postValues=%s\n", postValues)
		postDataStr := postValues.Encode()
		fmt.Printf("postDataStr=%s\n", postDataStr)
		postDataBytes := []byte(postDataStr)
		fmt.Printf("postDataBytes=%s\n", postDataBytes)
		postBytesReader := bytes.NewReader(postDataBytes)
		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("http get strUrl=%s response error=%s\n", strUrl, err.Error())
	}
	fmt.Printf("httpResp.Header=%s\n", httpResp.Header)
	fmt.Printf("httpResp.Status=%s\n", httpResp.Status)

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("get response for strUrl=%s got error=%s\n", strUrl, errReadAll.Error())
	}
	CurCookies = CurCookieJar.Cookies(httpReq.URL)
	respHtml = string(body)
	return respHtml
}

//get url response code
func getImg(strUrl string, postDict map[string]string) {

	httpClient := &http.Client{
		Jar: CurCookieJar,
	}

	var httpReq *http.Request
	if nil == postDict {
		httpReq, _ = http.NewRequest("GET", strUrl, nil)

	} else {
		postValues := url.Values{}
		for postKey, PostValue := range postDict {
			postValues.Set(postKey, PostValue)
		}
		postDataStr := postValues.Encode()
		postDataBytes := []byte(postDataStr)
		postBytesReader := bytes.NewReader(postDataBytes)
		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("http get strUrl=%s response error=%s\n", strUrl, err.Error())
	}

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("get response for strUrl=%s got error=%s\n", strUrl, errReadAll.Error())
	}
	CurCookies = CurCookieJar.Cookies(httpReq.URL)
	code, error := os.Create("D:/baiducode.png")
	if error != nil {
		fmt.Println(error)
	}
	code.Write([]byte(body))
	code.Close()
}

//打印cookie
func printCurCookies() {
	var cookieNum int = len(CurCookies)
	fmt.Printf("cookieNum=%d\r\n", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = CurCookies[i]
		fmt.Printf("curCk.Raw=%s\r\n", curCk.Value)
	}
}

//获取unix时间
func getMillisecond() int64 {
	MS := time.Now().Unix()
	return MS
}

// 密码加密
func RsaEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func main() {
	uname := ""
	upassword := ""

	p, err := phantomjs.Start()
	if err != nil {
		panic(err)
	}
	defer p.Exit() // Don't forget to kill phantomjs at some point.
	var gid interface{}
	var callback interface{}
	//获取gid
	err = p.Run("function gid(){ return 'xxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (e) { var t = 16 * Math.random() | 0,n = 'x' == e ? t : 3 & t | 8;return n.toString(16)}).toUpperCase()}", &gid)
	//获取callback()
	err = p.Run("function callback(){ return 'bd__cbs__'+Math.floor(2147483648 * Math.random()).toString(36)}", &callback)
	if err != nil {
		panic(err)
	}
	//获取token
	//tokenUrl := "https://passport.baidu.com/v2/api/?getapi&tpl=netdisk&subpro=netdisk_web&apiver=v3&tt="+strconv.FormatInt(getMillisecond(),10)+"&class=login&gid="+gid.(string)+"&logintype=basicLogin&callback="+callback.(string)
	//tokenUrl := "https://passport.baidu.com/v2/api/?getapi&tpl=netdisk&subpro=netdisk_web&apiver=v3&tt=1508470686&class=login&gid=E67758A-D8DB-45E0-B488-4DF52405EB73&logintype=basicLogin&callback=bd__cbs__lzwd5"
	//body1 := getUrlRespHtml(tokenUrl,nil)
	//fmt.Printf("第一次返回结果：%s\n",body1)
	//https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true
	var token string
	var getapiUrl string = "https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true"
	getApiRespHtml := getUrlRespHtml(getapiUrl, nil)

	loginTokenP, _ := regexp.Compile(`bdPass\.api\.params\.login_token='(?P<loginToken>\w+)';`)
	foundLoginToken := loginTokenP.FindStringSubmatch(getApiRespHtml)
	if nil != foundLoginToken {
		strLoginToken := foundLoginToken[1] //tmp go regexp not support named group, so use index here
		token = strLoginToken
	}
	//获取rsakey和pubkey

	rsaUrl := "https://passport.baidu.com/v2/getpublickey?token=" + token + "&tpl=netdisk&subpro=netdisk_web&apiver=v3&tt=" + strconv.FormatInt(getMillisecond(), 10) + "&gid=" + gid.(string) + "&callback=" + callback.(string)
	pubkeyHtml := getUrlRespHtml(rsaUrl, nil)
	println(pubkeyHtml)
	i := strings.Index(pubkeyHtml, "-----BEGIN PUBLIC KEY-----")
	j := strings.Index(pubkeyHtml, "-----END PUBLIC KEY-----")
	println("pubkey裁剪后：" + pubkeyHtml[i:j+26])
	str := pubkeyHtml[i : j+26]
	var pubkey string
	//pubkey = strings.Replace(str, "\n", "\\n", -1)
	pubkey = strings.Replace(str, "\\n", "\n", -1)
	pubkey2 := strings.Replace(pubkey, "\\/", "/", -1)
	println("pubkey替换后：" + pubkey)
	println("pubkey替换后2：" + pubkey2)
	k_i := strings.Index(pubkeyHtml, "\"key\":'")
	k_j := strings.LastIndex(pubkeyHtml, ",")
	key := pubkeyHtml[k_i+7 : k_j-1]
	println("key:" + key)
	//加密后的密码
	data, err := RsaEncrypt([]byte(pubkey2), []byte(upassword))
	pwd := base64.StdEncoding.EncodeToString(data)
	println(pwd)

	//准备post参数
	postDict := map[string]string{}
	postDict["staticpage"] = "https://passport.baidu.com/static/passpc-account/html/v3Jump.html"
	postDict["charset"] = "UTF-8"
	postDict["token"] = token
	postDict["tpl"] = "pp"
	postDict["subpro"] = ""
	postDict["apiver"] = "v3"
	postDict["tt"] = strconv.FormatInt(getMillisecond(), 10)
	postDict["codestring"] = ""
	postDict["safeflg"] = "0"
	postDict["u"] = "https://passport.baidu.com/"
	postDict["isPhone"] = ""
	postDict["detect"] = "1"
	postDict["gid"] = gid.(string)
	postDict["quick_user"] = "0"
	postDict["logintype"] = "basicLogin"
	postDict["logLoginType"] = "pc_loginBasic"
	postDict["idc"] = ""
	postDict["loginmerge"] = "true"
	postDict["username"] = uname
	postDict["password"] = pwd
	postDict["mem_pass"] = "on"
	postDict["rsakey"] = key
	postDict["crypttype"] = "12"
	postDict["ppui_logintime"] = "71755"
	postDict["countrycode"] = ""
	postDict["fp_uid"] = "a4cb1898d835565da9359337df7f1bc6"
	postDict["fp_info"] = "a4cb1898d835565da9359337df7f1bc6002~~~asaanRiuI~9is-0_haaFFiLBEF1B5GKBYX_iiLBEF1B5Gls-X_EapGqDapGqOaaqZapkZSikH~meF~VwFtc96xym4~GQE9qEIOF~suRgskR1tjqDBEB1rx6hPX__HhaqziuB-4iBX__rhaX-haXKhaqGig6bIO3t79HwqhFtsSPLT~NLHh0L6eFxIvPvujNvke3x99~x0JaX0nXGhYYXn9XO-EfVJ5XnQnmoZXGtX1a1XSXeXSXpGnOGXvZXpXlTL0yeISHgUuXrvOcOXXXZpZg70fyGG7fDmaGGNKIhaqVhaqxhaqLhaqXaaaMiu2xGV0Ol_uiA3LKj4vI-NLcdcL0_ahaqpassNfHBYAaaXehaXdas8TBnAZchaXJhaXUhaXlhaXmiErEnisYaysY0lCYajCY4OsjaOs0__"
	postDict["dv"] = "MDExAAoAvgALA2QAIwAAAF00AAwCACOJ2NjY2OhuOns1ciBhLHMsfC9_IBNME2MCcQJ1GmgMfA94HAcCAASRkZGRDAIAI4nY2NjY9xJGB0kOXB1QD1AAUwNcbzBvH34NfglmFHAAcwRgBwIABJGRkZEJAgAkiY2-v7i4uLi4nQMDVxZYH00MQR5BEUISTX4hfgt4HW8hQC1IBwIABJGRkZEIAgAhiYpAQX19fXC96ajmofOy_6D_r_ys88CfwLDRotGmybvfDQIAHZGRnOX9qeim4bPyv-C_77zss4DfgPCR4pHmifufBwIABJGRkZEJAgAkiY3Dwvv7-_v79sjInN2T1IbHitWK2onZhrXqtcWk16TTvM6qDQIAHZGRmfDovP2z9KbnqvWq-qn5ppXKleCT9oTKq8ajDQIABZGRkkBABwIABJGRkZEXAgAHkJMDAwF1FhYCACKwxK-fsYW8hbyIu4y6grOGv4e1hreGs4O7j7iKv4q5i72LBAIABpKSkJGkkgECAAaRk5ODju4FAgAEkZGRnRUCAAiRkZDP-AB_QxACAAGREwIAKJG1tbXdqd2t3uTL5JT1hvWF6pjswqDBqMy5l_Sb9tmvnbKN4Y7pgO4GAgAokZGR7e3t7e3t7evExMTGJiYmIGBgYGPn5-fhoaGhov7-_vioqKirxw0CAAWRkZHa2ggCACGJjaSlo6Ojq53JiMaB05LfgN-P3IzT4L_gleaD8b_es9YJAgAkiY2oqaysrKyspOvrv_6w96Xkqfap-ar6pZbJluOQ9YfJqMWgBwIABJGRkZEIAgAhiYpFRElJSUO14aDuqfu696j3p_Sk-8iXyL3Oq9mX9pv-DQIAHZGRnPTsuPm38KLjrvGu_q39opHOkeSX8oDOr8KnDAIAI4mBgYGBjpTAgc-I2pvWidaG1YXa6bbpmfiL-I_gkvaG9YLmCAIACZGXi4uXl5eHKAgCAB6Eh0ZGhYWFp2k9fDJ1J2YrdCt7KHgnFEsUcQNxHmwJAgAkiY2FhLu7u7u7k7q67q_hpvS1-Kf4qPur9MeYx7fWpdahzrzYBwIABJGRkZEMAgAjiejo6OjblsKDzYrYmdSL1ITXh9jrtOub-on6jeKQ9IT3gOQMAgAjievr6-vYJnIzfTpoKWQ7ZDRnN2hbBFsrSjlKPVIgRDRHMFQ"
	postDict["traceid"] = "2F0ADB01"
	var callback2 interface{}
	err = p.Run("function callback(){ return 'bd__cbs__'+Math.floor(2147483648 * Math.random()).toString(36)}", &callback2)
	postDict["callback"] = "parent." + callback2.(string)

	//第一次post
	onePost := "https://passport.baidu.com/v2/api/?login"

	body2 := getUrlRespHtml(onePost, postDict)
	println("omePost:" + body2)

	//获取codeString
	compile1, _ := regexp.Compile("codeString=(\\w+)&")
	match1 := compile1.FindString(body2)
	println("codestring:" + match1[11:len(match1)-1])
	codeString := match1[11 : len(match1)-1]
	postDict["codestring"] = codeString
	//获取验证码
	verifycodeUrl := "https://passport.baidu.com/cgi-bin/genimage?" + codeString
	getImg(verifycodeUrl, nil)
	//验证验证码
	var callback3 interface{}
	err = p.Run("function callback(){ return 'bd__cbs__'+Math.floor(2147483648 * Math.random()).toString(36)}", &callback3)
	var verifycode = ""
	println("请输入验证码:")
	fmt.Scanln(&verifycode)
	println("验证码：" + verifycode)
	l3, _ := url.Parse(verifycode)
	println("验证码：" + l3.Query().Encode())
	checkVerifycodeUrl := "https://passport.baidu.com/v2/?checkvcode&token=" + token + "&tpl=netdisk&subpro=netdisk_web&apiver=v3&tt=" + strconv.FormatInt(getMillisecond(), 10) + "&verifycode=" + verifycode + "&codestring=" + codeString + "&callback=" + callback3.(string)

	res3 := getUrlRespHtml(checkVerifycodeUrl, nil)
	println(res3)

	postDict["verifycode"] = verifycode
	postDict["ppui_logintime"] = "81755"

	lastUrl := "https://passport.baidu.com/v2/api/?login"
	res4 := getUrlRespHtml(lastUrl, postDict)
	println(res4)

}
