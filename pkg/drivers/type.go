package drivers

import "github.com/MontFerret/ferret/pkg/runtime/core"

var (
	HTTPHeaderType   = core.NewType("HTTPHeader")
	HTTPCookieType   = core.NewType("HTTPCookie")
	HTMLElementType  = core.NewType("HTMLElement")
	HTMLDocumentType = core.NewType("HTMLDocument")
)

// Comparison table of builtin types
var typeComparisonTable = map[core.Type]uint64{
	HTTPHeaderType:   0,
	HTTPCookieType:   1,
	HTMLElementType:  2,
	HTMLDocumentType: 3,
}

func Compare(first, second core.Type) int64 {
	f, ok := typeComparisonTable[first]

	if !ok {
		return -1
	}

	s, ok := typeComparisonTable[second]

	if !ok {
		return -1
	}

	if f == s {
		return 0
	}

	if f > s {
		return 1
	}

	return -1
}
