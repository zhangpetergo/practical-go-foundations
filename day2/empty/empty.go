package main

import (
	"fmt"
)

func main() {
	var i any
	// before 1.18
	// any == interface{}
	i = 7
	fmt.Println(i)
	i = "你好"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println(s)

	// switch type
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("unknown type %#v", i)
	}

	fmt.Println(maxInt([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxFloat([]float64{1, 2, 3, 4, 5}))
	fmt.Println(max([]int{1, 2, 3, 4, 5}))
	fmt.Println(max([]float64{1, 2, 3, 4, 5}))

}

// 泛型
// generics

type Number interface {
	int | float64
}

// wo can use that
// func max[T Number](nums []T) T {

// }

func max[T int | float64](nums []T) T {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func maxInt(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func maxFloat(floats []float64) float64 {
	if len(floats) == 0 {
		return 0
	}
	max := floats[0]
	for _, num := range floats {
		if num > max {
			max = num
		}
	}
	return max
}
