package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// bug 这是一个 closure
		//go func() {
		//	fmt.Println(i)
		//}()

		// fix 2 使用一个局部变量
		i := i
		go func() {
			fmt.Println(i)
		}()

		// fix 1 using a parameter
		//go func(n int) {
		//	fmt.Println(n)
		//}(i)
	}

	ch := make(chan string)
	// 程序运行到这里会阻塞，等待有人接收
	// 但是接收的程序片段在下面，由于无法到达下面，就形成死锁了
	//ch <- "hello"
	//msg := <-ch
	//fmt.Println(msg)

	go func() {
		ch <- "hello"
	}()
	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("Message: %#v", i+1)
			ch <- msg
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}
	msg = <-ch
	fmt.Printf("closed:%#v\n", msg)
	time.Sleep(10 * time.Millisecond)

	fmt.Println(sleepSort([]int{10, 48, 2, 5}))

}

func sleepSort(value []int) []int {
	ch := make(chan int)
	// 遍历slice
	for _, v := range value {
		// oh, a bug
		// forget
		v := v
		go func() {
			// 需要把int转换为duration类型
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()
	}
	var retVal []int
	for range value {
		n := <-ch
		retVal = append(retVal, n)
	}
	return retVal
}

func test() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("Message: %#v", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("hello", msg)
	}
}
