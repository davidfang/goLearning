package main

import (
	"context"
	"fmt"
	"time"
)

//withCancel取消控制
//https://segmentfault.com/a/1190000040917752
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
