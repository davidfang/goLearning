package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
	myPrint(s)
	s = append(s, 100, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
	// myPrint(s)

	a := []int{0, 1}
	b := [2]int{1, 2}
	fmt.Printf("原始 a: %v, %T\n", a, a)
	fmt.Printf("原始 b: %v, %T \n", b, b)
	changeSlice(a)
	changeArray(b)
	fmt.Printf("修改后 a: %v, %T\n", a, a)
	fmt.Printf("修改后 b: %v, %T \n", b, b)

	changeSlice(b[:])
	fmt.Printf("将数组转换成切片后修改 b: %v, %T \n", b, b)

}
//切片是引用传递，是可以改变值的
func changeSlice(s []int) {
	s[0] = 100
}
//数组是值传递，是不改变值的
func changeArray(a [2]int) {
	a[0] = 100
}

func myPrint(v []int) {
	fmt.Printf("%p len:%d, cap: %d \n", v, len(v), cap(v))
}
