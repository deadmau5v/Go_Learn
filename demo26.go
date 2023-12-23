package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}

		quit <- true
	}()

	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			println("退出...")
			return
		}
	}
}
