package main

import "time"

func main24() {
	// channel 和 range
	c := make(chan int, 100)

	go func() {
		for {
			c <- 1
			time.Sleep(1 * time.Second)
		}
	}()

	for i := range c {
		println(i)
	}
}
