package main

import (
	"fmt"
	"time"

	
)

func main() {
	// 1.timer基本使用
	// timer1 := time.NewTimer(2 * time.Second)
	// t1 := time.Now()
	// fmt.Printf("t1:%v\n", t1)
	// t2 := <-timer1.C
	// fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到了")	
	}
}