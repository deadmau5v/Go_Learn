package main

import (
	"runtime"
	"time"
)

func main25() {
	c := make(chan int, 3)
	c_ := make(chan int, 3)
	quit := make(chan bool, 1)

	go func() {
		c <- 1
		c_ <- 2

		c <- 3
		c_ <- 4

		c <- 5
		c_ <- 6

		c <- 7
		c_ <- 8

		quit <- true
	}()

	// 监控 Channel 状态
	for {
		select {

		case c1 := <-c:
			println(c1)
		case c1 := <-c_:
			println(c1)
		case <-quit:
			runtime.Goexit()
		case c_ <- 9:
			println(<-c_)
		default:
			println("默认")
		}
		time.Sleep(1 * time.Second)
	}

}
