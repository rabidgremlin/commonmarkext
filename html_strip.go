package commonmarkext

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

// dumb code to strip out html to populated alt tags
// as manner in which we process AST means we don't get to see raw text to store
// for alt tag

func recurseNodes(n *html.Node) string {

	buf := ""

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "alt" {
				buf += fmt.Sprint(a.Val)
				break
			}
		}
	}

	if n.Type == html.TextNode {
		buf += n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buf += recurseNodes(c)
	}

	return buf

}

//CMarkStripHTMLForAltTag stripts HTML from text, making it suitable for use in an alt tag.
//Not very pretty code at all.
func CMarkStripHTMLForAltTag(s string) string {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	r := recurseNodes(doc)

	return r
}
