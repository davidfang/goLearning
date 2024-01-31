package main

import (
	"fmt"
	"strings"
)

// func removeLastCity(str string) string {
// 	if str == "" {
// 		return ""
// 	}

// 	lastChar := str[len(str)-1]
// 	if lastChar == `市` {
// 		return str[:len(str)-1]
// 	}

// 	return str
// }

func removeCity(s string) string {
	if s == "" {
		return ""
	}
	index := strings.LastIndex(s, "市")
	if index == -1 {
		return s
	}
	return s[:index]
}

func main() {
	// fmt.Println(removeLastCity("北京市"))
	// fmt.Println(removeLastCity("上海"))
	// fmt.Println(removeLastCity(""))

	fmt.Println(removeCity("北京市"))
	fmt.Println(removeCity("上海"))
	fmt.Println(removeCity(""))
}
