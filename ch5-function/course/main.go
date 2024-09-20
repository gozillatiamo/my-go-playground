package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	err := chain("iambot")
	fmt.Println(err)
}

func chain(username string) error {
	err := take("firstname")
	return fmt.Errorf("username: %s got error: %v", username, err)
}

func take(name string) error {
	return fmt.Errorf("%s is wrong", name)
}

func fetchHtml(url string) (node *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}

	node, err = html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return node, err
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
