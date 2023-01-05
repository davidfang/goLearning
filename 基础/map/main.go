package main

import (
	"fmt"
	"sort"
)

func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("初始化之后")
	key := "中国"

	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")

	sliceMap[key] = value
	fmt.Println(sliceMap)

	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"

	fmt.Println(map1)
	sortMap(map1)
}

func sortMap(m map[int]string) {
	sli := []int{}
	for k, _ := range m {
		sli = append(sli, k)
	}

	sort.Ints(sli)
	fmt.Println(sli)
	for v := range sli {
		fmt.Println(v, sli[v], m[sli[v]])
	}
}
