package main

import "fmt"

func main() {
	//定义切片
	a := make([]int, 3)
	// value 0 0 0
	a = append(a, 11)
	// 添加11到切片中
	// value 0 0 0 11
	// index 0 1 2 3
	a[1] = 90
	// value 0 1 0 11
	fmt.Println(a)
	//数组的好处是什么?数组的好处是刚开始就定义了容量,不会再增加,可以确定使用的范畴,不浪费。
	//slice的好处是什么?我不确定这个数据量有多少,可以自动扩容。

}
