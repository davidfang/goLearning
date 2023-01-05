package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock() //加锁
		x = x + i
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
