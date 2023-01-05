package main

import (
	"fmt"
	"time"
)

func main() {
	select2()
}

func select1(){
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("i1从c1收到:%v \n",i1)
	case c2 <-i2:
		fmt.Printf("发送%v到c2 \n",i1)
	case i3,ok := <-c3:
		if ok {
			fmt.Printf("i3从c3收到:%v \n",i3)
		}else{
			fmt.Printf("c3关闭")
		}
	default:
		fmt.Println("什么也没有")
	}
}

func select2(){
	resChan := make(chan int)

	s := "abcd"
	for i ,n:= 0,len(s); i<n;i++{
		// resChan <- i
		fmt.Println(i,s[i])
	}

	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second *3):
		fmt.Println("请求超时!")
		// resChan <- 3
	//  default:
	// 	resChan <- 1
	}
}
func doData(data int){
	fmt.Printf("%v -----\n",data)
}