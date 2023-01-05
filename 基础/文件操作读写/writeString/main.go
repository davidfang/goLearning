package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./a.txt")
	if err != nil{
		fmt.Println("创建失败", err)
		os.Exit(1)
	}
	defer file.Close()

	if _,err :=io.WriteString(file,"./a.txt"); err != nil{
		fmt.Println("写入失败",err)
		os.Exit(1)
	}
}