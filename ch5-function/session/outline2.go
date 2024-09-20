package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	node, _ := fetchHtml(os.Args[1])
	forEachNode(node, startElement, endElement)
}

func fetchHtml(url string) (node *html.Node, err error) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("url: %s has is not OK", url)
	}
	node, err = html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("url: %s can't parse html %v", url, err)
	}
	return node, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}

}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
