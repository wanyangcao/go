package main

import "fmt"

// 空值返回0，非法字符串返回-1
func atoi(s string) int {
	if s == "" {
		return 0
	}
	isPlus := true
	switch s[0] {
	case '-':
		isPlus = false
		s = s[1:]
	case '+':
		isPlus = true
		s = s[1:]
	default:
		break
	}
	result := 0
	for _, v := range []byte(s) {
		if v < '0' || v > '9' {
			return -1
		}
		v -= '0'
		result = result*10 + int(v)
	}
	if !isPlus {
		result = -result
	}
	return result
}

func main() {
	a := "123a"
	fmt.Println(atoi(a))
}
