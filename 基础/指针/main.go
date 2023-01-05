package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	var a int
	fmt.Println(a)
	fmt.Println(&a)
	var p *int
	fmt.Println(p)
	p = &a
	*p = 20
	fmt.Println(a)
	fmt.Println(p)

	fmt.Printf("%v\t%v\t%v\t%v\t%v \n", KB, MB, GB, TB,PB)
	fmt.Printf("%T\t%T\t%T\t%T\t%T \n", KB, MB, GB, TB,PB)
}
