package main

import "fmt"

func main9() {
	func2("你好 ", "我叫", "dxj")
	var msg []string
	//msg[0] = "hello world"
	//msg[9] = "hello world"
	fmt.Printf("msg 类型为%T 长度为 %d \n", msg, len(msg))
}

// add 两个整数的和
func add(x, y int) int {
	return x + y
}

// func1 三个数的和 返回 float64
func func1(x, y int, f float32) float64 {
	return float64(x) + float64(y) + float64(f)
}

// func2 可变参数
func func2(args ...string) {
	fmt.Printf("类型为 %T\n", args)
	for i := range args {
		println(args[i])
	}
}
