package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("请输入内容:")
	fmt.Printf("请输入name:")
	fmt.Scanln(&name)
	fmt.Printf("请输入age:")
	fmt.Scanln(&age)
	fmt.Printf("请输入married:")
	fmt.Scanln(&married)
	// fmt.Scanf("%s %d %t",&name,&age,&married)
	// fmt.Scan(&name,&age,&married)

	fmt.Printf("接收结果： name:%s age:%d married:%t \n",name,age,married)
}