package metadata

import (
	"errors"
	"fmt"
)

var (
	// Normal ones
	CapLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LowLetters = []rune("abcdefghijklmnopqrstuvwxyz")
	// Math Bold Font
	MathBoldCap = []rune("𝐀𝐁𝐂𝐃𝐄𝐅𝐆𝐇𝐈𝐉𝐊𝐋𝐌𝐍𝐎𝐏𝐐𝐑𝐒𝐓𝐔𝐕𝐖𝐗𝐘𝐙")
	MathBoldLow = []rune("𝐚𝐛𝐜𝐝𝐞𝐟𝐠𝐡𝐢𝐣𝐤𝐥𝐦𝐧𝐨𝐩𝐪𝐫𝐬𝐭𝐮𝐯𝐰𝐱𝐲𝐳")

	// Map used for store the conversion relationship
	MyConvertMap = make(ConvertMap, 52) // 26 capital cases and 26 low cases
)

type Font int

const (
	MathBold Font = iota
)

type ConvertMap map[rune]rune

// For testing
func ShowConvertMap() {
	for k, v := range MyConvertMap {
		fmt.Printf("%s -> %s\n", string(k), string(v))
	}
}

// Choosing a font and initialize the convert map
func InitConvertMap(font Font) error {
	switch font {
	case MathBold:
		for i := 0; i < 26; i++ {
			MyConvertMap[CapLetters[i]] = MathBoldCap[i]
			MyConvertMap[LowLetters[i]] = MathBoldLow[i]
		}
		return nil
	default:
		return errors.New("the font is not supported")
	}
}
