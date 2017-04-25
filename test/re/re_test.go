package re

import (
	"fmt"
	"testing"
)

type Pkg struct {
	PkgName    string
	PreviewURL string
	OS         string
}

func TestR(t *testing.T) {
	uri := "http://cn.t.appcoachs.net/v3/win/WK7fQWfe/c3fe6b04298a11e789604fccf478c5c9/c7706fdc48f83c81feadf8d7f7076bd2aad8b72d?price=${AUCTION_PRICE}&currency=${AUCTION_CURRENCY}"
	R(uri)
	uri = "http://cn.t.appcoachs.net/v3/win/WK7fQWfe/c3fe6b04298a11e789604fccf478c5c9/c7706fdc48f83c81feadf8d7f7076bd2aad8b72d?currency=${AUCTION_CURRENCY}&price=${AUCTION_PRICE}"
	R(uri)
}
func TestCheckPkgName(t *testing.T) {
	pkg := Pkg{
		PkgName:    "382617920",
		PreviewURL: "https://itunes.apple.com/us/app/viber/id382617920?mt=8",
		OS:         "ios",
	}
	pkgName, cntModify := CheckPkgName(pkg.PkgName, pkg.PreviewURL, pkg.OS)
	fmt.Printf("Fixed pkgName is %s, changed times: %v\n", pkgName, cntModify)
	pkg = Pkg{
		PkgName:    "极品飞车19",
		PreviewURL: "https://itunesxapple.com/us/app/viber/id382617920?mt=8",
		OS:         "ios",
	}
	pkgName, cntModify = CheckPkgName(pkg.PkgName, pkg.PreviewURL, pkg.OS)
	ifFixed(pkgName, cntModify)
	pkg = Pkg{
		//PkgName:    "极品飞车19",
		PkgName:    "com.specialapps.datamanage",
		PreviewURL: "https://play.google.com/store/apps/details?id=com.specialapps.datamanager&hl=en",
		OS:         "android",
	}
	pkgName, cntModify = CheckPkgName(pkg.PkgName, pkg.PreviewURL, pkg.OS)
	ifFixed(pkgName, cntModify)
	pkg = Pkg{
		PkgName: "极品飞车19",
		//PkgName:    "com.specialapps.datamanage",
		PreviewURL: "https://play.google.com/store/apps/details?id=com.specialapps.datamanager&hl=en",
		OS:         "android",
	}
	pkgName, cntModify = CheckPkgName(pkg.PkgName, pkg.PreviewURL, pkg.OS)
	ifFixed(pkgName, cntModify)
	pkg = Pkg{
		PkgName: "极品飞车19",
		//PkgName:    "com.specialapps.datamanage",
		PreviewURL: "https://play.google.com/store/apps/details?hl=en",
		OS:         "android",
	}
	pkgName, cntModify = CheckPkgName(pkg.PkgName, pkg.PreviewURL, pkg.OS)
	ifFixed(pkgName, cntModify)
}

func ifFixed(pkgName string, cntModify int) {
	switch cntModify {
	case -1:
		fmt.Printf("Need to fixed, but pkgName is not found in preiview url: %s\n", pkgName)
	case 1:
		fmt.Printf("Fixed pkgName is %s, changed times: %v\n", pkgName, cntModify)
	default:
		fmt.Printf("Not need to fixed, pkgName: %s.\n", pkgName)

	}
}

func TestIsHostCN(t *testing.T) {
	hostname := "sd-cn01"
	isCN := isHostCN(hostname)
	fmt.Printf("is cn host %v.\n", isCN)
	hostname = "aws-CN01111111"
	isCN = isHostCN(hostname)
	fmt.Printf("is cn host %v.\n", isCN)
	hostname = "aws-us01"
	isCN = isHostCN(hostname)
	fmt.Printf("is cn host %v.\n", isCN)
}
