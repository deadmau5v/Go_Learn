package main

import "fmt"

func main4() {
	// 键盘的输入输出
	var (
		x int
		y float32
	)

	print("输入一个整数:\n>")
	fmt.Scanln(&x)
	print("输入一个浮点数:\n>")
	fmt.Scanln(&y)
	fmt.Printf("x = %d, y = %f", x, y)
	//fmt.Scan()
	//fmt.Scanf()
	//fmt.Scanln()
}
