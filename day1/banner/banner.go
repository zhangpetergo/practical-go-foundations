package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	fmt.Println(len("G☺"))

	s := "G☺"
	fmt.Printf("%c type is %T \n", s[0], s[0])

	for i := 0; i < len(s); i++ {

		fmt.Println(s[i])
	}
	for i, r := range s {
		fmt.Println(i, r)
		fmt.Printf("%c type is %T \n", r, r)
	}

	debug()
	fmt.Println("GOG isPalindRome?", isPalindRome("GOG"))
	fmt.Println("GO isPalindRome?", isPalindRomeB("GO"))
	fmt.Println("G☺G isPalindRome?", isPalindRomeB("G☺G"))
	fmt.Println("G☺ isPalindRome?", isPalindRomeB("G☺"))

}

// Unicode版本
func isPalindRomeB(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}

// 判断一个字符串是不是回文数
// 这个版本不能判断除了ASCII码之外的字符
// 普通版本
func isPalindRome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2  // Bug 是len是计算字符串的字节数，而不是字符的个数
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("_")
	}

	fmt.Println()

}

func debug() {
	x, y := 1, "1"
	fmt.Printf("x=%v , y=%v\n", x, y)
	fmt.Printf("x=%#v , y=%#v\n", x, y) // 在logging框架中使用%#v
}
