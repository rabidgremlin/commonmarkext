package commonmarkext

import (
	"bytes"
	"strings"
)

// pulled from go source and tweaked

type writer interface {
	WriteString(string) (int, error)
}

//const escapedChars = `&'<>"`
const escapedChars = `&<>"`

func escape(w writer, s string) error {
	i := strings.IndexAny(s, escapedChars)
	for i != -1 {
		if _, err := w.WriteString(s[:i]); err != nil {
			return err
		}
		var esc string
		switch s[i] {
		case '&':
			esc = "&amp;"
		//case '\'':
		// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
		//esc = "&#39;"
		case '<':
			esc = "&lt;"
		case '>':
			esc = "&gt;"
		case '"':
			// "&#34;" is shorter than "&quot;".
			esc = "&quot;"
		default:
			panic("unrecognized escape character")
		}
		s = s[i+1:]
		if _, err := w.WriteString(esc); err != nil {
			return err
		}
		i = strings.IndexAny(s, escapedChars)
	}
	_, err := w.WriteString(s)
	return err
}

//CMarkHTMLEscapeString escapes the supplied string using the commonmark HTML escaping rules.
func CMarkHTMLEscapeString(s string) string {
	if strings.IndexAny(s, escapedChars) == -1 {
		return s
	}
	var buf bytes.Buffer
	escape(&buf, s)
	return buf.String()
}
