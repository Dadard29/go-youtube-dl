package managers

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

var logger = log.NewLogger("MANAGER", logLevel.DEBUG)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func unidecode(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}
