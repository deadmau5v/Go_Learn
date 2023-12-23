package main

import (
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		println(i)
		time.Sleep(1 * time.Second)
	}
}

func main21() {
	defer println("Main Thread Exit...")
	go newTask()

	var res int
	go func(a, b int, res *int) {
		defer println("Main Thread Func Exit")
		defer runtime.Goexit()
		defer println("Main Thread Func Exit...")
		println("启动", a, b)
		*res = a + b
	}(10, 0, &res)

	time.Sleep(1 * time.Second)
	println("res: ", res)
	time.Sleep(2 * time.Second)

}
