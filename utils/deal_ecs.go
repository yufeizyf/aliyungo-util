package utils

import (
	"bufio"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"regexp"
	"strings"
)

func GetEcsDocs(module string) string {
	jsonstring, _ := simplejson.NewJson([]byte(module))

	version, _ := jsonstring.Get("ecs").Get("version").String()
	fmt.Println("ecs version: " + version)

	ecsList, _ := jsonstring.Get("ecs").Get("list").Array()

	//Get open-api
	var openAPI map[string]interface{}

	for _, d := range ecsList {
		element := d.(map[string]interface{})
		if element["name_en"] == "open-api" {
			openAPI = element
		}
	}
	oaFolder := openAPI["isFolder"].([]interface{})

	for _, d := range oaFolder {
		docs := d.(map[string]interface{}) //取出指定类型api信息，如instance，disk。。。
		name := docs["name_en"].(string)   //取出api名字
		apikey := docs["key"].(string)

		if ECSAPI[name] == true {
			docsFolder := docs["isFolder"].([]interface{})
			docsList := Funclist{}
			if name != "datatype" {
				for _, d := range docsFolder {
					element := d.(map[string]interface{})
					key := element["key"].(string) //只找pub的
					if strings.HasSuffix(key, "pub") {
						funcname := element["name_en"].(string)
						docsList[funcname] = true
						docFuncList[funcname] = apikey + " " + key + " " + name
					}
				}
			} else {
				for _, d := range docsFolder {
					element := d.(map[string]interface{})
					funcname := element["name_en"].(string)
					docsList[funcname] = true
				}
			}

			ecsDocs[name] = docsList
		}
	}

	return version //版本号用于区分ecs文档
}

func DealEcs(fileList []string, module string) {
	fmt.Println("Begin to add ecs docs links")

	GetEcsDocs(module)

	for i := 0; i < len(fileList); i++ {
		path := fileList[i]

		pkgPath, _, fileName := GetPackageAndFileName(path)

		fread, err := os.Open(path)
		defer fread.Close()

		if err != nil {
			fmt.Println(path, err)
			return
		} else {
			inbuff := bufio.NewReader(fread)

			newPath := NewFilePath(pkgPath, fileName)

			fwrite, err := os.Create(newPath)
			defer fwrite.Close()

			if err != nil {
				fmt.Println(newPath, err)
				return
			}

			for {
				line, err := inbuff.ReadString('\n')
				if err != nil || io.EOF == err {
					break
				}

				//^func
				isMatch, err := regexp.MatchString("^func", line)
				if isMatch {
					funcname := strings.ToLower(GetFuncName(line))

					apiName, isExist := isExistECsDocs(funcname)

					if isExist {
						urlDescribe := URL_ECS_PREFIX + apiName + "&" + funcname

						fwrite.WriteString("//\n")
						fwrite.WriteString("// You can read doc at ")
						fwrite.WriteString(urlDescribe)
						fwrite.WriteString("\n")
						fwrite.WriteString(line)

					} else {
						fwrite.WriteString(line)
					}

				} else if line == "//\n" || strings.HasPrefix(line, "// You can read doc at ") {
					// do nothing
				} else {
					fwrite.WriteString(line)
				}
			}

			WriteBackAndRemove(newPath, path)
		}
	}
	fmt.Println("finish!")
}

func isExistECsDocs(funcname string) (string, bool) {
	var apiName []string

	for key := range ecsDocs {
		apiName = append(apiName, key)
	}

	for k := range apiName {
		api := apiName[k]
		if ecsDocs[api][funcname] == true {
			return api, true
		}
	}
	return "", false
}

var docFuncList = make(map[string]string)

func DiffEcsDocAndApi(fileList []string, module string) (map[string]string, string) {
	fmt.Println("Begin to find those api in http://docs.aliyun.com but not in sdk")

	version := GetEcsDocs(module)

	var sdkFuncList = make(map[string]bool)

	var diffEcsDocAndSdkResult = make(map[string]string)

	for i := 0; i < len(fileList); i++ {
		path := fileList[i]

		fread, err := os.Open(path)
		defer fread.Close()

		if err != nil {
			fmt.Println(path, err)
			return nil, version
		} else {
			inbuff := bufio.NewReader(fread)

			for {
				line, err := inbuff.ReadString('\n')
				if err != nil || io.EOF == err {
					break
				}

				//^func
				isMatch, err := regexp.MatchString("^func", line)
				if isMatch {
					fname := strings.ToLower(GetFuncName(line))
					sdkFuncList[fname] = true
				}
			}
		}
	}

	keys := make([]string, 0, len(docFuncList))
	for k := range docFuncList {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		name := keys[i]

		if sdkFuncList[name] == false {
			diffEcsDocAndSdkResult[name] = docFuncList[name]
		}
	}

	fmt.Println("Docs has but not in SDK :")

	// for i := 0; i < len(diffEcsDocAndSdkResult); i++ {
	// 	fmt.Println(diffEcsDocAndSdkResult[i])
	// }

	for key, value := range diffEcsDocAndSdkResult {
		fmt.Println(key + "  " + value)
	}

	return diffEcsDocAndSdkResult, version
}
