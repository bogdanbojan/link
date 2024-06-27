package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// search is the struct that holds all the anchor hrefs and texts associated with them.
type search struct {
	href, text []string
}

func main() {
	file, err := os.ReadFile("ex2.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}

	s := search{}
	s.findAnchor(doc)
	s.printResults()
}

// findAnchor checks that the html node is an <a><\a> type and recurses through it.
// It calls findText on hit to append the text associated with the parent anchor node.
func (s *search) findAnchor(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		s.findText(n)
		for _, a := range n.Attr {
			s.href = append(s.href, a.Val)
			break
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.findAnchor(c)
	}
}

// findText recurses through the text nodes associated with the html Node passed in.
func (s *search) findText(n *html.Node) {
	if n.Type == html.TextNode {
		cleanText := strings.TrimSpace(n.Data)
		s.text = append(s.text, cleanText)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.findText(c)
	}
}

// printResults is a helper function that formats the response in a human friendy way.
func (s *search) printResults() {
	fmt.Println("RESULTS")
	fmt.Println("HREF: ", strings.Join(s.href, " "))
	fmt.Println("TEXT: ", strings.Join(s.text, " "))
}
