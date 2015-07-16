package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// This function add links for datatype of ecs.
func DealDataType(fileList []string, module string) {
	fmt.Println("Begin to add datatype docs links")

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
			} else {
				fmt.Println(newPath)
			}

			for {
				line, err := inbuff.ReadString('\n')

				if err != nil || io.EOF == err {
					break
				}

				//^func
				isMatch, err := regexp.MatchString("^type", line)
				if isMatch {
					dTypeName := strings.ToLower(GetDataType(line))
					if ecsDocs["datatype"][dTypeName] {
						urlDescribe := DATATYPE_URL_PREFIX + dTypeName

						fwrite.WriteString("//\n")
						fwrite.WriteString("// You can read doc at ")
						fwrite.WriteString(urlDescribe)
						fwrite.WriteString("\n")
						fwrite.WriteString(line)

					} else {
						fwrite.WriteString(line)
					}

				} else {
					fwrite.WriteString(line)
				}
			}

			WriteBackAndRemove(newPath, path)

		}
	}
	fmt.Println("finished!")
}

func GetDataType(line string) string {
	buff := []byte(line)

	var nameBytes []byte

	//跳过“type and   space”
	pos := 4
	for buff[pos] == ' ' {
		pos++
	}

	for i := pos; buff[i] != ' '; i++ {
		nameBytes = append(nameBytes, buff[i])
	}

	return string(nameBytes)
}
