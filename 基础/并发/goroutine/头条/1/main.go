package main

import "fmt"

//通过4个并发处理程序得出10以内的素数表

func main() {
	origin, wait := make(chan int), make(chan int)
	Processor(origin, wait)
	for num := 2; num < 100; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

func Processor(origin chan int, wait chan int) {
	go func() {
		prime, ok := <-origin
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range origin {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
