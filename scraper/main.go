package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type Result struct {
	url   string
	title string
}

func (r Result) String() string {
	return fmt.Sprintf("URL: %s \nTitle: %s", r.url, r.title)
}

func main() {
	urlToProcess := []string{
		"http://globo.com",
		"http://google.com",
		"https://globoesporte.globo.com/",
	}

	ini := time.Now()
	ch := make(chan Result, 3)
	for _, url := range urlToProcess {
		go scrap(url, ch)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("(Took ", time.Since(ini).Milliseconds(), "Milliseconds)")
}

func getTitle(html string) string {
	r, _ := regexp.Compile("<title>(.*?)<\\/title>")
	if title := r.FindStringSubmatch(html); len(title) > 0 {
		return r.FindStringSubmatch(html)[1]
	}
	return "Sem titulo "
}

func scrap(url string, ch chan Result) {
	r := Result{}
	fmt.Printf("try %s\n", url)
	r.url = url

	resp, _ := http.Get(url)
	html, _ := ioutil.ReadAll(resp.Body)
	r.title = getTitle(string(html))
	ch <- r
}
