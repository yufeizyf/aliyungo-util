package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
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

func Generate(funcName string, request []Parameter, response []Parameter) {
	GenerateRequest(funcName, request)
	GenerateResponse(funcName, response)
	GenerateFunc(funcName, request, response)
}

func GenerateRequest(funcName string, request []Parameter) {
	fwrite, err := os.OpenFile("/home/ubuntu/Documents/GoWork/src/github.com/denverdino/aliyungo/ecs/news.go", os.O_RDWR|os.O_APPEND, 0660)
	defer fwrite.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var line string

	line = "type " + funcName + "Args struct {\n"
	fwrite.WriteString(line)

	for i := 0; i < len(request); i++ {
		line = "\t" + request[i].pname + "\t" + request[i].ptype + "\n"
		fwrite.WriteString(line)
	}

	fwrite.WriteString("}\n\n")
}

func GenerateResponse(funcName string, response []Parameter) {
	fwrite, err := os.OpenFile("/home/ubuntu/Documents/GoWork/src/github.com/denverdino/aliyungo/ecs/news.go", os.O_RDWR|os.O_APPEND, 0660)
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
			line = "\t" + response[i].pname + "\t" + response[i].ptype + "\n"
			fwrite.WriteString(line)
		}
	}
	fwrite.WriteString("}\n\n")
}

func GenerateFunc(funcName string, request []Parameter, response []Parameter) {
	fwrite, err := os.OpenFile("/home/ubuntu/Documents/GoWork/src/github.com/denverdino/aliyungo/ecs/news.go", os.O_RDWR|os.O_APPEND, 0660)
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
	line = "\t" + "client.Invoke(\"" + funcName + "\", args, &response)\n"
	fwrite.WriteString(line)

	fwrite.WriteString("}\n\n")
}

func GenerateCodeTemplate(fileList []string, module string) {

	diffEcsDocAndSdkResult, version := DiffEcsDocAndApi(fileList, module)

	for _, value := range diffEcsDocAndSdkResult {
		funcInfo := strings.Split(value, " ") //funcInfo[0] apiKey  funcInfo[1] funcKey

		url := "https://docs.aliyun.com/getMd?url=cn/ecs/" + version + "/doc/0040-open-api/" + funcInfo[0] + "/" + funcInfo[1] + ".md"
		fmt.Println(url)
		markdown := GetMarkdown(url)
		funcName, request, response := ParseMarkDown(markdown)
		Generate(funcName, request, response)
	}
}
