package commonmarkext

import (
	//"net/url"
	"fmt"
	"strings"
)

// from https://gist.github.com/hnaohiro/4627658
// hacked to not encode chars < 127
func urlencode(s string) (result string) {
	for _, c := range s {

		if c <= 0x7f { // single byte
			result += string(c)
		} else if c > 0x1fffff { // quaternary byte
			result += fmt.Sprintf("%%%X%%%X%%%X%%%X",
				0xf0+((c&0x1c0000)>>18),
				0x80+((c&0x3f000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else if c > 0x7ff { // triple byte
			result += fmt.Sprintf("%%%X%%%X%%%X",
				0xe0+((c&0xf000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else { // double byte
			result += fmt.Sprintf("%%%X%%%X",
				0xc0+((c&0x7c0)>>6),
				0x80+(c&0x3f),
			)
		}
	}

	return result
}

//CMarkURLEncodeString escapes the supplied string using the commonmark HTML escaping rules.
//HACK HACK Commonmark encoding is weird. This code is dumb and not complete. Placeholder until we can get something better
func CMarkURLEncodeString(str string) string {

	str = strings.Replace(str, " ", "%20", -1)
	str = strings.Replace(str, "\"", "%22", -1)
	str = strings.Replace(str, "\\", "%5C", -1)
	str = strings.Replace(str, "ä", "%C3%A4", -1)
	str = strings.Replace(str, "[", "%5B", -1)
	str = strings.Replace(str, "]", "%5D", -1)
	str = strings.Replace(str, "`", "%60", -1)

	return urlencode(str)
}
