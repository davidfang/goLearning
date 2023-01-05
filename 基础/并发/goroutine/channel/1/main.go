package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//goroutine 将0-100发送到ch1
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//goroutine将从ch1中接收的值，将该值平方发送到ch2
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	//从ch2中接收值打印
	for i := range ch2 {
		fmt.Println(i)
	}
}