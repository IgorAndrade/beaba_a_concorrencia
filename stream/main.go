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
		"https://globoesporte.globo.com",
	}

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {

		ini := time.Now()
		ch := make(chan Result, 3)
		for _, url := range urlToProcess {
			go scrap(url, ch)
		}
		for i := 0; i < 3; i++ {
			r := <-ch
			//resp.Write([]byte(r.String()))
			fmt.Fprintln(resp, r)
			if f, ok := resp.(http.Flusher); ok {
				f.Flush()
			}
		}
		fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

	})
	http.ListenAndServe(":8083", nil)
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
