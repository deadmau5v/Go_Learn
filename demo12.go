package main

import "fmt"

func main12() {
	var arr []int = []int{4521, 4532, 3453, 4534, 355, 136}
	for i, value := range arr {
		fmt.Printf("index: %d value: %d\n", i, value)
	}

	var arr2 []int = arr[2:]
	arr2[0] = 10
	print("[")
	for _, value := range arr2 {
		print(value, ", ")
	}
	print("]\n")
	var arr3 = make([]int, 1000)
	arr3[len(arr3)-1] = 99999
	fmt.Printf("arr3 长度: %d, arr3 值: %d", len(arr3), arr3)

}
