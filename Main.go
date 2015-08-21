package main

import (
	"aliyungo-util/utils"
)

func main() {
	module := utils.ParseJsonData()

	filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

	fileList := utils.GetFileList(filePath)

	//utils.DealDataType(fileList, module)

	//utils.DealOss(fileList, module)

	utils.DealEcs(fileList, module)

	//utils.DiffEcsDocAndApi(fileList, module)

	//utils.GenerateCodeTemplate(filePath, fileList, module)
}
