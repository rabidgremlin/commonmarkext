package commonmarkext

import (
	"fmt"
	"regexp"
	"strings"
)

//HtmlGenerator is an example generator that generates commonmark compliant HTML.
type HtmlGenerator struct {
}

func NewHtmlGenerator() *HtmlGenerator {
	return new(HtmlGenerator)
}

func (h *HtmlGenerator) GenBlockQuote(content string) string {
	return fmt.Sprintf("<blockquote>%s</blockquote>\n", content)
}

func (h *HtmlGenerator) GenList(numberedList bool, listStartNumber int, content string) string {

	if numberedList {
		if listStartNumber == 1 {
			return fmt.Sprintf("<ol>%s</ol>\n", content)
		} else {
			return fmt.Sprintf("<ol start=\"%d\">%s</ol>\n", listStartNumber, content)
		}

	} else {
		return fmt.Sprintf("<ul>%s</ul>\n", content)
	}

}

func (h *HtmlGenerator) GenItem(content string) string {
	return fmt.Sprintf("<li>%s</li>\n", content)
}

func (h *HtmlGenerator) GenCodeBlock(fenceInfo string, content string) string {

	words := strings.Split(fenceInfo, " ")

	if words[0] != "" {
		return fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>\n", words[0], CMarkHTMLEscapeString(content))
	} else {
		return fmt.Sprintf("<pre><code>%s</code></pre>\n", CMarkHTMLEscapeString(content))
	}
}

func (h *HtmlGenerator) GenHTML(content string) string {
	return fmt.Sprintf("%s", content)
}

func (h *HtmlGenerator) GenParagraph(inTightList bool, content string) string {

	if inTightList {
		return content
	} else {
		return fmt.Sprintf("<p>%s</p>\n", content)
	}
}

func (h *HtmlGenerator) GenHeader(level int, content string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, content, level)
}

func (h *HtmlGenerator) GenHrule() string {
	return "<hr />\n"
}

func (h *HtmlGenerator) GenText(content string) string {
	return CMarkHTMLEscapeString(content)
}

func (h *HtmlGenerator) GenSoftbreak() string {
	return "\n"
}

func (h *HtmlGenerator) GenLinebreak() string {
	return "<br />\n"
}

func (h *HtmlGenerator) GenCode(content string) string {
	return fmt.Sprintf("<code>%s</code>", CMarkHTMLEscapeString(content))
}

func (h *HtmlGenerator) GenInlineHTML(content string) string {
	return fmt.Sprintf("%s", content)
}

func (h *HtmlGenerator) GenEmph(content string) string {
	return fmt.Sprintf("<em>%s</em>", content)
}

func (h *HtmlGenerator) GenStrong(content string) string {
	return fmt.Sprintf("<strong>%s</strong>", content)
}

func (h *HtmlGenerator) GenLink(linkUrl, title, content string) string {

	encodedUrl := CMarkURLEncodeString(linkUrl)
	encodedUrl = CMarkHTMLEscapeString(encodedUrl)

	if title != "" {
		return fmt.Sprintf("<a href=\"%s\" title=\"%s\">%s</a>", encodedUrl, CMarkHTMLEscapeString(title), content)
	} else {
		return fmt.Sprintf("<a href=\"%s\">%s</a>", encodedUrl, content)
	}
}

func (h *HtmlGenerator) GenImage(imgUrl, alt, title string) string {

	encodedUrl := CMarkURLEncodeString(imgUrl)
	encodedUrl = CMarkHTMLEscapeString(encodedUrl)

	if title != "" {
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\" title=\"%s\" />", encodedUrl, CMarkStripHTMLForAltTag(alt), CMarkHTMLEscapeString(title))
	} else {
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />", encodedUrl, CMarkStripHTMLForAltTag(alt))
	}
}

func (h *HtmlGenerator) GenDocument(content string) string {
	content = strings.Replace(content, "><li", ">\n<li", -1)
	content = strings.Replace(content, "><ul", ">\n<ul", -1)
	content = strings.Replace(content, "><ol", ">\n<ol", -1)
	content = strings.Replace(content, "><p", ">\n<p", -1)
	content = strings.Replace(content, "><hr", ">\n<hr", -1)
	content = strings.Replace(content, "><h", ">\n<h", -1)
	content = strings.Replace(content, "<blockquote></blockquote>", "<blockquote>\n</blockquote>", -1)
	content = strings.Replace(content, "<blockquote><blockquote>", "<blockquote>\n<blockquote>", -1)

	re := regexp.MustCompile("(.+)<blockquote>")
	content = re.ReplaceAllString(content, "$1\n<blockquote>")

	re = regexp.MustCompile("(.+)<ul>")
	content = re.ReplaceAllString(content, "$1\n<ul>")

	re = regexp.MustCompile("(.+)<ol>")
	content = re.ReplaceAllString(content, "$1\n<ol>")

	return content
}
