package re

import "regexp"
import "fmt"
import "strings"
import "net/url"

func CheckPkgName(pkgName, previewLink, os string) (string, int) {
	var cntModify int
	if os == "ios" {
		re := regexp.MustCompile("^[0-9]+$")
		if !re.MatchString(pkgName) {
			re = regexp.MustCompile("id([0-9]+)")
			reLink := regexp.MustCompile(`itunes\.apple\.com`)
			if re.MatchString(previewLink) && reLink.MatchString(previewLink) {
				// capture mathed content in ()
				pkgName = re.FindAllStringSubmatch(previewLink, -1)[0][1]
				cntModify = 1
			} else {
				cntModify = -1
			}

		}
	} else if os == "android" {
		re := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
		isPkgName := false
		pkgNameVals := strings.Split(pkgName, ".")
		fmt.Printf("is pkgName or not: %v, pkgNameVals %v.\n", isPkgName, pkgNameVals)
		for _, val := range pkgNameVals {
			isPkgName = re.MatchString(val)
		}
		fmt.Printf("Android mathed %v, isPkgName %v\n", pkgName, isPkgName)
		// Fix pkgName error format
		if !isPkgName {
			uri, _ := url.Parse(previewLink)
			reLink := regexp.MustCompile(`play\.google\.com`)
			if len(uri.Query().Get("id")) > 0 && reLink.MatchString(previewLink) {
				pkgName = uri.Query().Get("id")
				cntModify = 1
			} else {
				cntModify = -1
			}
		}
	}
	return pkgName, cntModify
}

func isHostCN(hostname string) bool {
	isCN := regexp.MustCompile(`(\-[cC][nN][0-9]+)`).MatchString(hostname)
	return isCN
}

func R(uri string) {
	//uri := "http://cn.t.appcoachs.net/v3/win/WK7fQWfe/c3fe6b04298a11e789604fccf478c5c9/c7706fdc48f83c81feadf8d7f7076bd2aad8b72d?price=${AUCTION_PRICE}&currency=${AUCTION_CURRENCY}"
	fmt.Printf("%v\n", uri)
	//uri = regexp.MustCompile(`&currency=\${AUCTION_CURRENCY}|currency=\${AUCTION_CURRENCY}&`).ReplaceAllString(uri, "")
	uri = regexp.MustCompile(`(|&)currency=\${AUCTION_CURRENCY}(&|)`).ReplaceAllString(uri, "")
	fmt.Printf("%v\n", uri)
}
