package util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"strings"
)

// This function parse json data from http://docs.aliyun.com.
func ParseJsonData() string {
	var module string

	doc, err := goquery.NewDocument("http://docs.aliyun.com")
	if err != nil {
		fmt.Print(err)
	}

	doc.Find("script").Each(func(i int, contentSelection *goquery.Selection) {
		content := contentSelection.Text()

		cons := strings.Split(content, ";")

		for i := 0; i < len(cons); i++ {
			if strings.Contains(cons[i], "window.docModule") {
				module = cons[i]
				break
			}
		}
	})

	module = strings.TrimPrefix(module, "\nwindow.docModule=JSON.parse('")
	module = strings.TrimSuffix(module, "')")

	return module
}
