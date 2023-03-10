package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)//启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait()//等待所有登记的goroutine都结束
}

func hello(i int) {
	defer wg.Done()//goroutine 结束就登记-1
	fmt.Println("Hello Goroutine!",i)
}