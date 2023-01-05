package main

import (
	"fmt"
	"time"
)

//生产者:生成factor 整数倍的序列
func Producer(factor int,out chan<- int) {
	for i := 0;; i++{
		out <- i * factor
	}
}

//消费者
func Consumer(in <-chan int){
	for v:=range in{
		fmt.Println(v)
	}
}

func main()  {
	ch := make(chan int,64)//成果队列
	go Producer(3,ch)//生成3的倍数的序列
	go Producer(5,ch)//生成5的倍数的序列
	go Consumer(ch)//消费生成的序列

	//运行一定时间后退出
	time.Sleep(5 * time.Second)

}