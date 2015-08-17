package main

import (
	"aliyungo-util/utils"
)

func main() {
	module := utils.ParseJsonData()

	fileList := utils.GetFileList("/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo")

	//utils.DealDataType(fileList, module)

	//utils.DealOss(fileList, module)

	//utils.DealEcs(fileList, module)

	//utils.DiffEcsDocAndApi(fileList, module)

	utils.GenerateCodeTemplate(fileList, module)

	//funcName, request, response := utils.ParseMarkDown(utils.GetMarkdown("https://docs.aliyun.com/getMd?url=cn%2Fecs%2F0.2.2%2Fdoc%2F0040-open-api%2F0060%25E9%2595%259C%25E5%2583%258F%25E7%259B%25B8%25E5%2585%25B3%25E6%258E%25A5%25E5%258F%25A3%2F0021%25E4%25BF%25AE%25E6%2594%25B9%25E9%2595%259C%25E5%2583%258F%25E5%25B1%259E%25E6%2580%25A7pub.md"))

	//utils.Generate(funcName, request, response)
}
