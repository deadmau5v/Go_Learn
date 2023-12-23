package main

import "fmt"

type type16_1 struct {
	name string
	age  int
}

func main16() {
	func16_1("hello, world", 123, "hello", type16_1{name: "hello", age: 19})

	var str1 interface{} = "hello"
	str1 = 1999
	value, ok := str1.(string)
	println(value, ok)
}

func func16_1(args ...interface{}) interface{} {
	var _, ok = args[0].(string)
	if !ok {
		fmt.Println()
		return false
	} else {
		fmt.Println(args)
		return true
	}
}
