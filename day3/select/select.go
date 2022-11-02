package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	// 使用上下文来做超时
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	// 初步理解：每次接受一个，哪个先到选择哪个
	select {
	case val := <-ch1:
		fmt.Println("ch1:", val)
	case val := <-ch2:
		fmt.Println("ch2:", val)
	//超时
	//case <-time.After(5 * time.Millisecond):
	//	fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println("timeout")
	}

}
