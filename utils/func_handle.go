package utils

import ()

// This function gets function names from xxx.go file.
func GetFuncName(line string) string {
	buff := []byte(line)

	var nameBytes []byte

	pos := 4 //跳过“func”
	for buff[pos] == ' ' {
		pos++
	}

	if buff[pos] != '(' { // 处理类型  func (a A) funcname(){}
		for i := pos; buff[i] != '('; i++ {
			nameBytes = append(nameBytes, buff[i])
		}
	} else { // 处理类型  func funcname(){}
		i := pos
		for buff[i] != ')' {
			i++
		}

		i = i + 2 // ") " 跳过空格

		for buff[i] != '(' {
			nameBytes = append(nameBytes, buff[i])
			i++
		}
	}

	return string(nameBytes)
}
