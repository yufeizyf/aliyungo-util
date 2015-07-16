package main

import (
	"aliyungo-util/utils"
)

func main() {
	module := utils.ParseJsonData()

	fileList := utils.GetFileList("/home/ubuntu/Documents/GoWork/src/github.com/denverdino/aliyungo")

	//utils.DealDataType(fileList, module)

	//utils.DealOss(fileList, module)

	//utils.DealEcs(fileList, module)

	utils.DiffEcsDocAndApi(fileList, module)
}
