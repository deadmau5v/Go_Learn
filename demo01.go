package main

import (
	"fmt"
	"strconv"
)

var globalName = "helloworld"
var test = "version"

func main1() {
	// 调用fmt模块 打印
	var test = "run"
	fmt.Println("hello world")

	// 注释
	// 为开头 单行注释

	/*
		多行注释
	*/

	// 变量
	// var 变量名 类型名 = 值
	var name string = "包皮哥"
	var age int = 18
	var 傻逼 string = "你妈"
	var ____1 int = 1

	fmt.Println("我是"+name, "刚满"+strconv.Itoa(age)+"岁")
	fmt.Println(傻逼, ____1)

	var (
		name2           string
		age2            int
		addr            string
		password        int
		username, email string = "None", "None"
		hhh                    = "感觉不如var()"
	)
	hhhh := "自动推导类型"

	fmt.Println(name2, age2, addr, password, username, email, hhhh, hhh)

	// 几种打印 print println printf
	print("hello world\n")
	println("hello", "world")
	fmt.Println("hello", "world"+"!")
	fmt.Printf("%s\n", "hello world")
	println("hello world")

	// 取地址符号 &
	fmt.Printf("value: [%d], memory of: [%p]\n", age, &age)
	age = 200
	fmt.Printf("value: [%d], memory of: [%p]\n", age, &age)

	println("========")
	println(age, &age, age2, &age2)
	age, age2 = age2, age
	println(age, &age, age2, &age2)
	println("========")

	// 匿名变量
	var i = func() (a int, b int) {
		return 1, 2
	}

	var a, _ = i()
	println(a)

	// 作用域
	var f1 = func() {

	}

	var f2 = func() {

	}
	var localName = "local"
	print(f1, f2)
	print(globalName, localName)
	println(test)
	// 全局变量

	const URL string = "www.baidu.com"
	const URL2 string = "fanyi.baidu.com"
	const (
		URL3       string = "hello"
		URL4       string = "hello4"
		URL5, URL6        = "nmsl", "cnm"
	)

	const (
		c0 = "hello" //0
		c1 = iota    // 1
		c2
		c3
		c4 = "hello"
		c5
		c6
		c7
		c8 = iota //8
	)

	const (
		c01 = iota //0
		c02        //1
	)

	const c001 = iota //0
	const (
		c0001 = iota //0
	)

}
