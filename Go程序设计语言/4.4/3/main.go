package main

import "fmt"

type Point struct{ X, Y int }
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{}
	w2 := Wheel{Circle: Circle{Point: Point{8, 8}, Radius: 5}, Spokes: 20}
	fmt.Printf("w1: %v\n", w1)
	fmt.Printf("w2: %v\n", w2)
	fmt.Printf("w1#: %#v\n", w1)
	fmt.Printf("w2#: %#v\n", w2)

}