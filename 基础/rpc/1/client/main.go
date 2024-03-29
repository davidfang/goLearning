package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//参数
type Params struct {
	Width, Height int
}

func main() {
	//1.连接远程RPC服务
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8100")
	if err != nil {
		log.Fatal(err)
	}
	//2.调用方法
	//面积
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)
	//周长
	per := 0
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &per)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", per)
}
