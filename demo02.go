package main

import (
	"fmt"
)

func main3() {
	// 数据类型

	// bool 类型
	var isFlag bool = true
	var notFlag bool = !isFlag
	println(notFlag)
	fmt.Printf("%T %t %t\n", notFlag, notFlag, "hello")

	// 整形 浮点型
	var i1, i2 int = 1, 2
	var f1, f2 float32 = 1.0, 2.0
	var ff1, ff2 float64 = 3.14159261234561111, 3.14159267473585
	fmt.Printf("%T %.10\n", ff1, ff2)
	fmt.Printf("%T %.10f\n", f1, f2)
	fmt.Printf("%T %d\n", i1, i2)
	var ui1 uint = 100
	var s1 int = 1
	var s2 int8 = 8
	var s3 int16 = 16
	var s4 int32 = 32
	var s5 int64 = 64
	println(ui1, s1, s2, s3, s4, s5)
	var b1 byte = 255
	println(b1)
	var r1 rune = 10 ^ 5
	println(r1)
	var f01 float32 = 0.3
	var f02 float32 = 0.2
	println(f01 + f02)
	println()

	// 字符串
	var str1 string = "String"
	var str2 string = "Hello world"
	fmt.Printf("%T %s\n", str1, str2)
	println(str1 + str2)
	println("\"转义字符\"\"\n\t\b\a")

	// 类型转换
	var int1 int = 1000
	var float1 float32 = float32(int1)
	fmt.Printf("%T %f", float1, float1)
	//fmt.Printf("%T %f", str3, str3)

	// 运算符
	var int2 int = 123
	int2++
	println(int2)

	println(*&int2)

	var x1 int = 10
	var address *int = &x1
	println(*address)
}
