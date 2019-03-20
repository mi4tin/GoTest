package main

import (
	"fmt"
	"moqikaka.com/goutil/webUtil"
	"regexp"
	"testing"
)

func TestGetTitleByUrl(t *testing.T) {
	res, err := webUtil.GetWebData("http://www.bejson.com/", nil)
	if err != nil {
		fmt.Println("err1", err)
		return
	}
	//fmt.Println("resData", string(res))
	resData := string(res)
	fmt.Println("title", getPreTag(resData))
}

var preTagReg = regexp.MustCompile(`<title>(.+)</title>`)

func getPreTag(tag string) string {
	matchs := preTagReg.FindStringSubmatch(tag)
	if len(matchs) == 2 {
		return matchs[1]
	}
	return ""
}
