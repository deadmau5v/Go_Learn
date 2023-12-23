package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main23() {
	var channel1 = make(chan int, 100)
	go func() {
		for {
			channel1 <- rand.Int() * 1000.0
			time.Sleep(1 * time.Second)
			if rand.Int()%14 == 0 {
				close(channel1)
				return
			}
		}

	}()

	var reduce int
	for {
		time.Sleep(1 * time.Second)
		if num, ok := <-channel1; ok {
			reduce += num
			fmt.Printf("%d\n", reduce)
		} else {
			break
		}
	}
	println("程序结束")

}
