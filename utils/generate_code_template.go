package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Parameter struct {
	pname string // parameter name
	ptype string // parameter type
}

func GetMarkdown(link string) *goquery.Document {

	encode, _ := url.Parse(link)

	markdown, err := goquery.NewDocument("https://docs.aliyun.com/getMd?" + encode.Query().Encode())
	if err != nil {
		fmt.Print(err)
	}

	return markdown
}

func format(s string) string {
	s = strings.TrimLeft(s, "\r\n")

	return strings.TrimRight(s, "\r\n")
}

func ParseMarkDown(markdown *goquery.Document) (string, []Parameter, []Parameter) {
	var funcName string
	var request = make([]Parameter, 0)
	var response = make([]Parameter, 0)

	markdown.Find("table").Each(func(ta int, tableSelection *goquery.Selection) {
		if tableSelection.HasClass("ecs ecs_interface_request") {
			tableSelection.Find("tr").Each(func(tr int, trSelection *goquery.Selection) {
				p := Parameter{}

				isPara := false
				trSelection.Find("td").Each(func(td int, tdSelection *goquery.Selection) {
					isPara = true

					if tr == 1 {
						if td == 3 {
							fun := strings.Split(tdSelection.Text(), "：")
							funcName = format(fun[len(fun)-1])
						}
					} else {
						if td == 0 {
							p.pname = format(tdSelection.Text())
						}
						if td == 1 {
							p.ptype = format(tdSelection.Text())
						}
					}
				})

				if isPara == true && tr != 1 {
					request = append(request, p)
				}
				isPara = false
			})
		}

		if tableSelection.HasClass("ecs ecs_interface_response") {
			tableSelection.Find("tr").Each(func(tr int, trSelection *goquery.Selection) {
				p := Parameter{}

				isPara := false
				trSelection.Find("td").Each(func(td int, tdSelection *goquery.Selection) {
					isPara = true

					if td == 0 {
						p.pname = format(tdSelection.Text())
					}
					if td == 1 {
						p.ptype = format(tdSelection.Text())
					}
				})

				if isPara == true {
					response = append(response, p)
				}
			})
		}
	})

	fmt.Println("api name: ", funcName)
	fmt.Println("request: ", request)
	fmt.Println("response: ", response)

	return funcName, request, response
}

func Generate(funcName string, request []Parameter, response []Parameter, apiType string, path string) {
	file_temp := path + string(filepath.Separator) + apiType + ".go"
	fmt.Println(file_temp)

	if FileExist(file_temp) == false {
		os.Create(file_temp)
	}

	GenerateRequest(funcName, request, file_temp)
	GenerateResponse(funcName, response, file_temp)
	GenerateFunc(funcName, request, response, file_temp)
}

func IsInteger(str string) string {
	if str == "Integer" {
		return "int"
	}

	return str
}

func GenerateRequest(funcName string, request []Parameter, filename string) {
	fwrite, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
	defer fwrite.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var line string

	line = "type " + funcName + "Args struct {\n"
	fwrite.WriteString(line)

	for i := 0; i < len(request); i++ {
		line = "\t" + request[i].pname + "\t" + strings.ToLower(IsInteger(request[i].ptype)) + "\n"
		fwrite.WriteString(line)
	}

	fwrite.WriteString("}\n\n")
}

func GenerateResponse(funcName string, response []Parameter, filename string) {
	fwrite, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
	defer fwrite.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var line string

	line = "type " + funcName + "Response struct {\n"
	fwrite.WriteString(line)

	if len(response) == 0 {
		fwrite.WriteString("\tCommonResponse\n")
	} else {
		fwrite.WriteString("\tCommonResponse\n")

		for i := 0; i < len(response); i++ {
			line = "\t" + response[i].pname + "\t" + strings.ToLower(IsInteger(response[i].ptype)) + "\n"
			fwrite.WriteString(line)
		}
	}
	fwrite.WriteString("}\n\n")
}

func GenerateFunc(funcName string, request []Parameter, response []Parameter, filename string) {
	fwrite, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
	defer fwrite.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var line string

	line = "func (client *Client) " + funcName + "(" + "args *" + funcName + "Args)"

	if len(response) == 0 {
		line = line + " error {\n"
		fwrite.WriteString(line)
	} else {
		line = line + " (*" + funcName + "Response, error) {\n"
		fwrite.WriteString(line)
	}

	//func body
	//response := DescribeInstanceMonitorDataResponse{}
	line = "\t" + "response := " + funcName + "Response{}\n"
	fwrite.WriteString(line)

	//err = client.Invoke("DescribeInstanceMonitorData", args, &response)
	line = "\t" + "client.Invoke(\"" + funcName + "\", &args, &response)\n"
	fwrite.WriteString(line)

	fwrite.WriteString("}\n\n")
}

func GenerateEcsTempFolder(path string) string {
	err := os.Mkdir(path+string(filepath.Separator)+"ecs_temp", 0777)

	if err != nil {
		fmt.Println(err, "creating ecs__temp folder failed")
		return ""
	}

	return path + string(filepath.Separator) + "ecs_temp"
}

func GenerateCodeTemplate(path string, fileList []string, module string) {
	ecs_temp := GenerateEcsTempFolder(path)

	diffEcsDocAndSdkResult, version := DiffEcsDocAndApi(fileList, module)

	fmt.Println(diffEcsDocAndSdkResult)

	for _, value := range diffEcsDocAndSdkResult {
		funcInfo := strings.Split(value, " ") //funcInfo[0] apiKey  funcInfo[1] funcKey  funcInfo[2] 函数所属类型(instance, disk....)

		url := "https://docs.aliyun.com/getMd?url=cn/ecs/" + version + "/doc/0040-open-api/" + funcInfo[0] + "/" + funcInfo[1] + ".md"
		fmt.Println(url)
		markdown := GetMarkdown(url)
		funcName, request, response := ParseMarkDown(markdown)
		Generate(funcName, request, response, funcInfo[2], ecs_temp)
	}
}
