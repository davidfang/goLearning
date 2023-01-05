package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d \n",c,n)
			c = n
		}
	}

	var a = []int{1, 3, 4, 5, 6}
    fmt.Printf("slice a : %v , len(a) : %v , cap(a) : %v\n", a, len(a), cap(a))
    b := a[1:2]
    fmt.Printf("slice b : %v , len(b) : %v , cap(b) : %v\n", b, len(b), cap(b))
    d := b[0:3]
    fmt.Printf("slice d : %v , len(d) : %v , cap(d) : %v\n", d, len(d), cap(d))


	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    d1 := slice[6:8]
    fmt.Println(d1, len(d1), cap(d1))
    d2 := slice[1:6:8]
    fmt.Println(d2, len(d2), cap(d2))
}