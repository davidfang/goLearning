package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for{
		line,err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完成")
				break
			}else{
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}
	


}