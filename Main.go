package main

import (
	"aliyungo-util/utils"
)

func main() {
	//module := utils.ParseJsonData()

	//fileList := utils.GetFileList("/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo")

	//utils.DealDataType(fileList, module)

	//utils.DealOss(fileList, module)

	//utils.DealEcs(fileList, module)

	//utils.DiffEcsDocAndApi(fileList, module)

	//utils.GenerateCodeTemplate(fileList, module)
	utils.ParseMarkDown(utils.GetMarkdown("https://docs.aliyun.com/getMd?url=cn/ecs/0.2.2/doc/0040-open-api/0040磁盘相关接口/0100扩容磁盘pub.md"))
}
