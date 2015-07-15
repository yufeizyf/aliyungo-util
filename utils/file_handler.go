package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// This function traverse a specified file path to get all *.go files, not including *test.go files
func GetFileList(path string) []string {
	var fileList []string

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		goFile := strings.HasSuffix(path, ".go")
		if goFile {
			testFile := strings.HasSuffix(path, "[*]test.go")
			if testFile == false {
				fileList = append(fileList, path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Errorf("filepath.Walk() returned %v\n", err)
		return nil
	}

	return fileList
}

//	This function gets package path, package name and file name of a go file from its absolute path
func GetPackageAndFileName(path string) (string, string, string) {
	directory := strings.Split(path, "/")

	fileName := directory[len(directory)-1]

	pkgPath := getPackagePath(path, fileName)

	pkgName := directory[len(directory)-2]

	return pkgPath, pkgName, fileName
}

// This function returns package path. For example, /a/b/c/ecs/ is package path for /a/b/c/ecs/example.go
func getPackagePath(path string, filename string) string {
	result := strings.Split(path, filename)

	return result[0]
}

// This function will return a temp file path for a go file. For example, /a/example.go ==> /a/example_temp.go
func NewFilePath(pkgPath string, fileName string) string {
	oldName := strings.Split(fileName, ".")[0]

	newName := oldName + "_temp.go"

	newFilePath := pkgPath + newName

	return newFilePath
}

// This function write content from xxx_temp.go to xxx.go, then delete xxx_temp.go.
func WriteBackAndRemove(temp string, src string) {
	fread, err := os.Open(temp)
	defer fread.Close()

	if err != nil {
		fmt.Errorf(temp, err)
		return
	} else {
		inbuff := bufio.NewReader(fread)

		fwrite, err := os.Create(src)
		defer fwrite.Close()

		if err != nil {
			fmt.Errorf(src, err)
			return
		}

		for {
			line, err := inbuff.ReadString('\n')

			if err != nil || io.EOF == err {
				fmt.Errorf("Writing file occurs some problem")
				break
			} else {
				fwrite.WriteString(line)
			}
		}

		os.Remove(temp)
	}
}
