package main

import (
	"fmt"
	"sync"
	"time"
)

// 声明全局等待组
var wg sync.WaitGroup

// 声明全局锁
var mutex sync.Mutex

// 声明全局余票
var ticket int = 10

func main() {
	// 设置等待组计数器
	wg.Add(3)
	// 窗口卖票
	go saleTicket2("窗口A", &wg)
	go saleTicket2("窗口B", &wg)
	go saleTicket2("窗口C", &wg)
	wg.Wait()
	fmt.Println("运行结束!")
}

// 卖票流程
func saleTicket(windowName string, wg *sync.WaitGroup) {
	//卖票流程结束后关闭
	defer wg.Done()
	for {
		//加锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(10 * time.Millisecond)
			ticket--
			fmt.Printf("%s 卖出一张票，余票：%d \n", windowName, ticket)
		} else {
			fmt.Printf("%s 票已经售完！\n", windowName)
			//解锁
			mutex.Unlock()
			break

		}
		//解锁
		mutex.Unlock()
	}
}

// 卖票流程
func saleTicket2(windowName string, wg *sync.WaitGroup) {
	// 卖票流程结束后关闭
	for {
		// 加锁
		if ticket > 0 {
			time.Sleep(10 * time.Millisecond)
			ticket--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, ticket)
		} else {
			fmt.Printf("%s 票已卖完! \n", windowName)
			// 解锁
			break
		}
		// 解锁
	}
}
