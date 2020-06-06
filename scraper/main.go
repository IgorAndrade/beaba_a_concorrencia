package main

import (
	"fmt"
	"regexp"
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
	fmt.Println(urlToProcess)
}

func getTitle(html string) string {
	r, _ := regexp.Compile("<title>(.*?)<\\/title>")
	if title := r.FindStringSubmatch(html); len(title) > 0 {
		return title[1]
	}
	return "Sem titulo "
}
