package utils

import (
	"bufio"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"regexp"
)

// This function get oss docs which exists in http://docs.aliyun.com
func GetOssDocs(module string) {
	jsonstring, _ := simplejson.NewJson([]byte(module))

	version, _ := jsonstring.Get("oss").Get("version").String()
	fmt.Println("now oss version is : " + version)

	ossList, _ := jsonstring.Get("oss").Get("list").Array()

	//Get oss api-reference
	var apiRefer map[string]interface{}

	for _, d := range ossList {
		element := d.(map[string]interface{})
		if element["name_en"] == "api-reference" {
			apiRefer = element
		}
	}

	arFolder := apiRefer["isFolder"].([]interface{})

	for _, d := range arFolder {
		docs := d.(map[string]interface{})
		name := docs["name_en"].(string) //api-reference

		if OSSAPI[name] == true {
			docsFolder := docs["isFolder"].([]interface{})
			docsList := Funclist{}

			for _, d := range docsFolder {
				element := d.(map[string]interface{})
				apiname := element["name_en"].(string)
				docsList[apiname] = true
			}
			ossDocs[name] = docsList
		}
	}
}

// Describe this function.
func DealOss(fileList []string, module string) {
	fmt.Println("Begin to add oss docs links")

	GetOssDocs(module)

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
			} else {
				fmt.Println(newPath)
			}

			for {
				line, err := inbuff.ReadString('\n')
				if err != nil || io.EOF == err {
					break
				}

				//^func
				isMatch, err := regexp.MatchString("^func", line)
				if isMatch {
					funcname := GetFuncName(line)

					apiName, isExist := isExistOssDocs(funcname)

					if isExist {
						urlDescribe := URL_OSS_PREFIX + apiName + "&" + ossChart[funcname]

						fwrite.WriteString("//\n")
						fwrite.WriteString("// You can read doc at ")
						fwrite.WriteString(urlDescribe)
						fwrite.WriteString("\n")
						fwrite.WriteString(line)

					} else {
						fmt.Println("not exist ", funcname)
						fwrite.WriteString(line)
					}

				} else {
					fwrite.WriteString(line)
				}
			}

			WriteBackAndRemove(newPath, path)
		}
	}
	fmt.Println("finish!")
}

func isExistOssDocs(funcname string) (string, bool) {
	var apiName []string

	for key := range ossDocs {
		apiName = append(apiName, key)
	}

	for k := range apiName {
		api := apiName[k]

		if ossDocs[api][ossChart[funcname]] == true {
			return api, true
		}
	}
	return "", false
}
