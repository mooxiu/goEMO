package app

import (
	"fmt"
	"goemo/metadata"
)

func ConvertEmo(wholeText, ident string, startPoint int, m metadata.ConvertMap) string {
	// string to []rune{}
	wholeTextRune := []rune(wholeText)
	// TODO: when ident is not pure ASCII
	endPoint := len(ident) + startPoint

	beforeIdent := string(wholeTextRune[:startPoint])
	afterIdent := string(wholeTextRune[endPoint:])

	replacedIdent := make([]rune, 0)
	for _, v := range wholeTextRune[startPoint : endPoint+1] {
		mv := m[v]
		replacedIdent = append(replacedIdent, mv)
	}

	return fmt.Sprint(beforeIdent, string(replacedIdent), afterIdent)
}
