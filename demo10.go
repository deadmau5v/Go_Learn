package main

// 别名
import format "fmt"

func main10() {
	// 值传递
	var arr []int = []int{1, 2, 3, 4}
	format.Println(arr)
	// 引用传递
	arr = update_array(arr, 2, 5)
	format.Println(arr)

}

func update_array(arr []int, index, value int) []int {
	var array []int
	copy(array, arr)
	for i := 0; i < len(arr); i++ {
		if i == index {
			array = append(array, value)
		} else {
			array = append(array, arr[i])
		}
	}
	return array
}

func init() {
	println("初始化... demo10")
}
