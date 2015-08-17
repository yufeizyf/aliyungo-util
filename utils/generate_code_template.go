package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetMarkdown(url string) *goquery.Document {

	markdown, err := goquery.NewDocument("https://docs.aliyun.com/getMd?url=cn/ecs/0.2.2/doc/0040-open-api/0030实例相关接口/0010创建实例pub.md")
	if err != nil {
		fmt.Print(err)
	}

	return markdown
}

func ParseMarkDown(markdown *goquery.Document) {
	fmt.Println(markdown)
}

func GenerateCodeTemplate(fileList []string, module string) {

	diffEcsDocAndSdkResult, version := DiffEcsDocAndApi(fileList, module)

	for _, value := range diffEcsDocAndSdkResult {
		funcInfo := strings.Split(value, " ") //funcInfo[0] apiKey  funcInfo[1] funcKey

		url := "https://docs.aliyun.com/getMd?url=cn/ecs/" + version + "/doc/0040-open-api/" + funcInfo[0] + "/" + funcInfo[1] + ".md"
		fmt.Println(url)
		markdown := GetMarkdown(url)
		ParseMarkDown(markdown)
	}
}
