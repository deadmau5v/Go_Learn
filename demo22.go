package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main22() {
	// channel
	c := make(chan int, 30)

	go func() {
		for {
			// 发送一个数据 到管道
			c <- int(rand.Float32() * 10.0)
			time.Sleep(time.Duration(int(rand.Float32()*100.0) * int(time.Millisecond)))
		}
	}()

	var num int
	for {
		num = <-c
		if num == num {
			fmt.Printf("管道长度: %d 元素数量: %d 占用 %.2f%s \n", len(c), cap(c), float32(len(c))/float32(cap(c))*100, "%")
		}
		time.Sleep(time.Duration(int(rand.Float32()*150.0) * int(time.Millisecond)))

	}
}
