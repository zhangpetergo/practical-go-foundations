package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	// log的输出是并发安全的
	// 我们可以查看log的源码
	//siteTime("https://www.baidu.com")
	urls := []string{
		"https://baidu.com",
		"https://google.com",
		"https://no-such-site.biz",
	}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		url := url
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	wg.Wait()
}

// 访问一个网址的时间
func siteTime(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s->%s", url, err)
		return
	}

	// 如果上面有错误，应该返回，否则这里为nil
	defer resp.Body.Close()
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("ERROR: %s->%s", url, err)
	}
	duration := time.Since(start)
	log.Printf("INFO: %s->%v", url, duration)
}
