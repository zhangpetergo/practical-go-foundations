package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                // a slice of int
	fmt.Println("len", len(s)) // len is "nil" safe

	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2) // fmt的格式化输出  %#v

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v\n", s3) // half open range 左闭右开

	fmt.Printf("s2: len(%d) cap(%d)\n", len(s2), cap(s2))

	s4 := make([]int, 4, 6)
	fmt.Printf("s4: len(%d) cap(%d)\n", len(s4), cap(s4))

	s4 = s4[:5]
	fmt.Printf("s4: len(%d) cap(%d)\n", len(s4), cap(s4))

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))

	values := []float64{2, 1, 3, 4}
	values2 := []float64{2, 1, 3}
	fmt.Println(median(values))
	fmt.Println(median(values2))
	fmt.Println(median(nil))
}

// 连接两个slice
func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2)+1)
	copy(s[:len(s1)], s1)
	copy(s[len(s1):], s2)
	// s3 := append(s1, s2...)
	return s
}

// 找到slice的中位数
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	// fmt.Printf("%p\n", &values)
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	n := 2
	i := len(nums) / n
	if len(nums)%2 == 1 {
		return nums[i], nil
	} else {
		return (nums[i-1] + nums[i]) / 2, nil
	}
	// fmt.Printf("%p\n", &values)
	// fmt.Printf("%p\n", &nums)
}
