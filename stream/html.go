package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Result struct {
	url   string
	title string
}

func (r Result) String() string {
	return fmt.Sprintf("URL: %s \nTitle: %s", r.url, r.title)
}

func hasClass(attribs []html.Attribute, className string) bool {
	for _, attr := range attribs {
		if attr.Key == "class" && strings.Contains(attr.Val, className) {
			return true
		}
	}
	return false
}

func getFirstTextNode(htmlParsed *html.Node) *html.Node {
	if htmlParsed == nil {
		return nil
	}

	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Type == html.TextNode {
			return m
		}
		r := getFirstTextNode(m)
		if r != nil {
			return r
		}
	}
	return nil
}

func getFirstElementByClass(htmlParsed *html.Node, elm, className string) *html.Node {
	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Data == elm && hasClass(m.Attr, className) {
			return m
		}
		r := getFirstElementByClass(m, elm, className)
		if r != nil {
			return r
		}
	}
	return nil
}

func getFirstElement(htmlParsed *html.Node, elm string) *html.Node {
	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Data == elm {
			return m
		}
		r := getFirstElement(m, elm)
		if r != nil {
			return r
		}
	}
	return nil
}

func scrap(url string, ch chan Result) {
	r := Result{}
	fmt.Printf("try %s\n", url)
	r.url = url

	resp, _ := http.Get(url)
	body := resp.Body
	htmlParsed, _ := html.Parse(body)

	h1 := getFirstTextNode(getFirstElement(htmlParsed, "title"))
	if h1 != nil {
		r.title = strings.TrimSpace(h1.Data)
	} else {
		fmt.Println("Scrap error: Can't find title. url:'", url, "'")
	}
	ch <- r
}