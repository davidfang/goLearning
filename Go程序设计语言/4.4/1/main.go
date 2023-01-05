package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//添加树
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//将元素按照顺序追加到values里面，然后返回结果slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func main() {
	intA := []int{109, 24, 3, 50}
	fmt.Printf("原始：%v \n", intA)
	Sort(intA)
	fmt.Printf("排序后：%v \n",intA)

}
