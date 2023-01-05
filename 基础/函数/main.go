package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
func main() {
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%+v", o)

	fmt.Printf("加和：%v \n", plus(1, 2, 3, 3, 4, 3))

	//闭包
	c := a()
	c()
	c()
	a() //不输出

	c2 := a()
	c2()
	c2()

	tmp1 := b(100)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := b(1000)
	fmt.Println(tmp2(1), tmp2(2))

	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))

	//递归
	var i int = 7
	//阶乘
	fmt.Println("阶乘：")
	fmt.Printf("%d  = %d \n",i, factorial(i))
	//斐波那契数列
	fmt.Println("斐波那契数列:")
	for i=0;i<10;i++{
		fmt.Printf("%-5v %d \n",i,fibonaci(i))
	}

	//defer
	ts := []Test{{"a"}, {"b"}, {"c"}, {"d"},}
	for _,t := range ts{
		defer Close(t)
	}

	//lock
	func() {
        t1 := time.Now()

        for i := 0; i < 1000000; i++ {
            testLock()
        }
        elapsed := time.Since(t1)
        fmt.Println("test elapsed: ", elapsed)
    }()
    func() {
        t1 := time.Now()

        for i := 0; i < 1000000; i++ {
            testLockdefer()
        }
        elapsed := time.Since(t1)
        fmt.Println("testdefer elapsed: ", elapsed)
    }()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int  = make(chan int, 10)
	close(ch)
	ch <-5
}
//斐波那契数列
func fibonaci(i int) int {
	if i == 0{
		return 0
	}
	if i == 1{
		return 1
	}
	//这个数列从第3项开始，每一项都等于前两项之和
	return fibonaci(i-1)+fibonaci(i-2)
}
//阶乘
func factorial(i int) int {
	if i <= 1 {
		return i
	}
	return i * factorial(i-1)
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}
func myfunc(args ...interface{}) {

}

func plus(args ...int) (r int) {
	for _, i := range args {
		r += i
	}
	return r
}

//闭包
func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Printf("(%v)闭包：%#v  %5v  \n", &i, i, i)
		return i
	}
	return b
}

// 外部引用函数参数局部变量
func b(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " 关闭")
}

func Close(t Test){
	t.Close()
}

func testLock() {
    lock.Lock()
    lock.Unlock()
}

func testLockdefer() {
    lock.Lock()
    defer lock.Unlock()
}