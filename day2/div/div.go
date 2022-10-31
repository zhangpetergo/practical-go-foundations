package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(safeDiv(1, 0))
	fmt.Println("hello")
}

// 有名字的返回值
// named rerurn value
func safeDiv(a, b int) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, nil
}
func div(a, b int) int {
	return a / b
}
