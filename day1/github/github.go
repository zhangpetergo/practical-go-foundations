package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10809")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get("https://api.github.com/users/rowerdog")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy %s", err)
	// }

	var r Reply
	de := json.NewDecoder(resp.Body)
	if err := de.Decode(&r); err != nil {
		log.Fatalf("decode failed - %s", err)
	}
	fmt.Println(r)

	name, num, _ := githubInfo("rowerdog")
	fmt.Println(name, num)
}

// github Info returns name,number of repos for login
func githubInfo(login string) (string, int, error) {
	// first , we request with url

	// url.PathEscape �?我们的字符能在url�?显示出来
	urls := "htts://api.github.com/users/" + url.PathEscape(login)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:10809")
	}
	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(urls)
	if err != nil {
		// log.Fatalf("error: %s", err)
		return "", 0, err
	}
	// 匿名结构�?
	var r struct {
		Login        string `json:"login,omitempty"`
		Public_Repos int    `json:"public___repos,omitempty"`
	}
	decode := json.NewDecoder(resp.Body)
	if err := decode.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Login, r.Public_Repos, nil
}

type Reply struct {
	Login        string `json:"login,omitempty"`
	Public_Repos int    `json:"public___repos,omitempty"`
}
