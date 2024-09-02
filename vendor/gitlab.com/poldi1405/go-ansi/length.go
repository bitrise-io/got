package ansi

import (
	"regexp"
	"unicode/utf8"
)

func GetLengthWithoutCodes(content string) int {
	var re = regexp.MustCompile(`(?m)\x1b\[([\d;]+m|[\d;]+(H|A|B|C|D|J|K|S|T)|s|u)`)
	return utf8.RuneCountInString(re.ReplaceAllString(content, ""))
}
