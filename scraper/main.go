package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/tsuru/tsuru/log"
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
	in := time.Now()
	ch := make(chan Result, 3)
	for _, url := range urlToProcess {
		go scrap(url, ch)
	}

	for i := 0; i < 3; i++ {
		r := <-ch
		fmt.Println(r)
	}
	fmt.Println("levou ", time.Since(in).Milliseconds())
	//time.Sleep(time.Second)
}

func scrap(url string, ch chan Result) {
	r := Result{url: url}
	fmt.Println("tentando url: %s", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.title = getTitle(string(html))

	ch <- r

}

func getTitle(html string) string {
	r, _ := regexp.Compile("<title>(.*?)<\\/title>")
	if title := r.FindStringSubmatch(html); len(title) > 0 {
		return title[1]
	}
	return "Sem titulo "
}
