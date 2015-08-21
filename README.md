# aliyungo-util
addlink的重构版  
为阿里云的Go SDK 的ecs和oss添加在 http://docs.aliyun.com/ 对应的文档链接 现有功能:
> 1 添加ecs的api和datatype对应的docs链接  
> 2 添加oss的api对应的docs链接   
> 3 查看docs.aliyun.com存在，sdk不存在的api  
> 4 基于功能３实现代码模板生成

# 用法
1. add docs links for ecs.

		func main() {
			module := utils.ParseJsonData()

			filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

			fileList := utils.GetFileList(filePath)

			utils.DealEcs(fileList, module)
			
			utils.DealDataType(fileList, module)
		}
		
2. add docs links for oss.

		func main() {
			module := utils.ParseJsonData()

			filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

			fileList := utils.GetFileList(filePath)

			utils.DealOss(fileList, module)
		}
		
3. add docs links for datatype.

		func main() {
			module := utils.ParseJsonData()

			filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

			fileList := utils.GetFileList(filePath)
			
			utils.DealDataType(fileList, module)
		}
		
4. find api in docs.aliyun.com but not in sdk.

		func main() {
			module := utils.ParseJsonData()

			filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

			fileList := utils.GetFileList(filePath)

			utils.DiffEcsDocAndApi(fileList, module)
		}
5. generate code template.

		func main() {
			module := utils.ParseJsonData()

			filePath := "/Users/zhangyf/Documents/GoWork/src/github.com/denverdino/aliyungo"

			fileList := utils.GetFileList(filePath)

			utils.GenerateCodeTemplate(filePath, fileList, module)
		}