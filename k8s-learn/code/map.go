package main

import "fmt"

// map
func main() {
	//golang是强类型语言,所有数据都必须定义类型,除了interface以外,这个interface是定义所有的数据类型。
	msg := map[string]interface{}{"username": "admin", "password": "12345"}
	users := map[string]string{"username": "root", "username1": "admin"}
	fmt.Println(msg, users)
	//map 和struct有什么区别
	//Order := struct {
	//	Uid  string
	//	Type string
	//	time string
	//}{}

	//struct的使用场景?固定好了的数据结构,不可以变更。
	//map的使用场景?不用固定数据类型,可以随便增加减少。
	stu1 := new(stu) // 初始化结构体 new的过程可以理解为面向对象的实例化的过程!
	stu1.name = "admin"
	stu1.id = 1
	fmt.Println(stu1)
}

type stu struct { ////type xxx struct 固定结构 这里的xxx是你的struct的名字
	name string // k v k是什么？定义的数据结构的名字，v是
	id   int
}
