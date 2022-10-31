package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("day2/freq/sherlock-sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	// 打开文件等，不要忘了关闭连接
	defer file.Close()

	// 这种字符串在处理时因为转义会出现问题
	//var path = "C:\to\new\report.csv"

	// 我们使用 raw string
	var path = `C:\to\new\report.csv`
	fmt.Println(path)

	// 我们可以使用原生字符串创建多行

	//var request = `Get http1.1
	//Host: golang.org
	//Connection: Close
	//`

	freqs, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("error: ", err)
	}
	word, err := maxWord(freqs)
	if err != nil {
		log.Fatalf("error: ", err)
	}
	fmt.Println("max time word", word)
}

// 设置正则表达式
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func mostCommon(r io.Reader) (string, error) {
	frequency, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(frequency)
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}
	maxN, maxW := 0, ""
	for k, v := range freqs {
		if v > maxN {
			maxN = v
			maxW = k
		}
	}
	return maxW, nil
}

// return word's frequency
func wordFrequency(r io.Reader) (map[string]int, error) {

	scan := bufio.NewScanner(r)
	freqs := make(map[string]int, 0)
	//lnum := 0
	for scan.Scan() {
		words := wordRe.FindAllString(scan.Text(), -1)
		for _, v := range words {
			freqs[strings.ToLower(v)]++
		}
		//lnum++
		//scan.Text()
	}
	if err := scan.Err(); err != nil {
		log.Fatalf("error: %s", err)
	}
	//fmt.Printf("lines: %d\n", lnum)
	return freqs, nil
}
