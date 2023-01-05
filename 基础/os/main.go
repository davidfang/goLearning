package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil{
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("文件大小：",size)

	// 定义接收文件读取的字节数组
	buf:=bufio.NewReader(file)


	for {
		line, err := buf.ReadString('\n')
		if err != nil{
			if err == io.EOF {
				fmt.Println("文件读取结束")
				break
			}else{
				fmt.Println("文件读取错误：",err)
				return
			}
		}
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		field := strings.Fields(line)
		// fmt.Println(field)
		switch field[0] {	
		case "mkdir":
			fmt.Println("创建文件夹：",field[1])
			os.Mkdir(field[1],os.ModePerm)
		case "cp":
			// fmt.Println("复制文件：",field[1]," 到 ",field[2])
			copyFiel(field[1],field[2])
		}
	}
}

func copyFiel(from, to string) {
	//打开源文件
	file1,err1 := os.Open(from)
	if err1 != nil{
		fmt.Println(from,"文件打开错误",err1)
		return
	}
	defer file1.Close()
	//创建目标文件
	file2,err2 := os.OpenFile(to,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err2 != nil{
		fmt.Println(to,"文件创建失败",err2)
	}
	defer file2.Close()
	n,e :=io.Copy(file2,file1)
	if e!=nil{
		fmt.Println( "复制失败:",from,to,e)
	}else{
		fmt.Println(to,"复制成功， 文件大小:",n/1024/1024,"Mb")
	}

}
