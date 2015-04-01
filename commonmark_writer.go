package commonmarkext

import (
	"fmt"
	"github.com/rhinoman/go-commonmark"
	"log"
)

//CommonmarkGenerator is the interface needs to be implemented by every output generator.
//Each of these methods will be called as the particular commonmark element is found in the input document.
type CommonmarkGenerator interface {
	// blocks
	GenBlockQuote(content string) string
	GenList(numberedList bool, listStartNumber int, content string) string
	GenItem(content string) string
	GenCodeBlock(fenceInfo string, content string) string
	GenHTML(content string) string
	GenParagraph(inTightList bool, content string) string
	GenHeader(level int, content string) string
	GenHrule() string

	// inlines
	GenText(content string) string
	GenSoftbreak() string
	GenLinebreak() string
	GenCode(content string) string
	GenInlineHTML(content string) string
	GenEmph(content string) string
	GenStrong(content string) string
	GenLink(linkURL, title, content string) string
	GenImage(imgURL, alt, title string) string

	// helpers
	GenDocument(content string) string
}

//GenerateFromString generates output for the supplied string (containing commonmark text), using the supplied generator.
func GenerateFromString(commonmarkString string, generator CommonmarkGenerator) string {
	doc := commonmark.ParseDocument(commonmarkString, commonmark.CMARK_OPT_DEFAULT)
	defer doc.Free()

	fmt.Println(doc.RenderXML(commonmark.CMARK_OPT_DEFAULT))

	return GenerateFromNode(doc, generator)
}

//GenerateFromBytes generates output for the supplied byte array (containing commonmark text), using the supplied generator.
func GenerateFromBytes(commonmarkBytes []byte, generator CommonmarkGenerator) []byte {
	inStr := string(commonmarkBytes[:])

	doc := commonmark.ParseDocument(inStr, commonmark.CMARK_OPT_DEFAULT)
	defer doc.Free()

	outStr := GenerateFromNode(doc, generator)

	return []byte(outStr)
}

//GenerateFromNode generates output for the supplied commonmark node using the supplied generator.
func GenerateFromNode(node *commonmark.CMarkNode, generator CommonmarkGenerator) string {

	if node.GetNodeType() == commonmark.CMARK_NODE_NONE {
		return ""
	}

	content := ""

	nn := node.FirstChild()

	for {
		if nn.GetNodeType() == commonmark.CMARK_NODE_NONE {
			break
		}

		content += GenerateFromNode(nn, generator)
		nn = nn.Next()
	}

	switch node.GetNodeType() {
	case commonmark.CMARK_NODE_DOCUMENT:
		return generator.GenDocument(content)
	case commonmark.CMARK_NODE_BLOCK_QUOTE:
		return generator.GenBlockQuote(content)
	case commonmark.CMARK_NODE_LIST:
		return generator.GenList(node.GetListType() == commonmark.CMARK_ORDERED_LIST, node.GetListStart(), content)
	case commonmark.CMARK_NODE_ITEM:
		return generator.GenItem(content)
	case commonmark.CMARK_NODE_CODE_BLOCK:
		return generator.GenCodeBlock(node.GetFenceInfo(), node.GetLiteral())
	case commonmark.CMARK_NODE_HTML:
		return generator.GenHTML(node.GetLiteral())
	case commonmark.CMARK_NODE_PARAGRAPH:
		parentsParent := node.Parent().Parent()

		if parentsParent.GetNodeType() == commonmark.CMARK_NODE_LIST && parentsParent.GetListTight() {
			return generator.GenParagraph(true, content)
		}
		return generator.GenParagraph(false, content)
	case commonmark.CMARK_NODE_HEADER:
		return generator.GenHeader(node.GetHeaderLevel(), content)
	case commonmark.CMARK_NODE_HRULE:
		return generator.GenHrule()

	//Inline
	case commonmark.CMARK_NODE_TEXT:
		return generator.GenText(node.GetLiteral())
	case commonmark.CMARK_NODE_SOFTBREAK:
		return generator.GenSoftbreak()
	case commonmark.CMARK_NODE_LINEBREAK:
		return generator.GenLinebreak()
	case commonmark.CMARK_NODE_CODE:
		return generator.GenCode(node.GetLiteral())
	case commonmark.CMARK_NODE_INLINE_HTML:
		return generator.GenInlineHTML(node.GetLiteral())
	case commonmark.CMARK_NODE_EMPH:
		return generator.GenEmph(content)
	case commonmark.CMARK_NODE_STRONG:
		return generator.GenStrong(content)
	case commonmark.CMARK_NODE_LINK:
		return generator.GenLink(node.GetUrl(), node.GetTitle(), content)
	case commonmark.CMARK_NODE_IMAGE:
		return generator.GenImage(node.GetUrl(), content, node.GetTitle())
	default:
		log.Panicf("Unexpected node type: %s  content: %s", node.GetNodeTypeString(), content)
		return ""
	}

}
