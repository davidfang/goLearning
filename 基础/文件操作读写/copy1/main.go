package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./a.txt")
	if err != nil{
		fmt.Println("读取失败", err)
		os.Exit(1)
	}
	defer file.Close()

	if _,err :=io.Copy(os.Stdout,file); err != nil{
		fmt.Println("复制输出失败", err)
		os.Exit(1)
	}
}