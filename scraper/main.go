package main

import (
	"fmt"
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
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	ini := time.Now()
	ch := make(chan Result, 3)
	for _, url := range urlToProcess {
		go scrap(url, ch)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}
