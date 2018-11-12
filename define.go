package fion

type LogStyle int

//FD: Foreground Default
//FB: Foreground Bright
//BD: Background Default
//BB: Background Bright
const (
	SFDBlack LogStyle = iota
	SFDRed
	SFDGreen
	SFDYellow
	SFDBlue
	SFDMagenta
	SFDCyan
	SFDWhite
	SFBBlack
	SFBRed
	SFBGreen
	SFBYellow
	SFBBlue
	SFBMagenta
	SFBCyan
	SFBWhite
	SBDBlack
	SBDRed
	SBDGreen
	SBDYellow
	SBDBlue
	SBDMagenta
	SBDCyan
	SBDWhite
	SBBBlack
	SBBRed
	SBBGreen
	SBBYellow
	SBBBlue
	SBBMagenta
	SBBCyan
	SBBWhite
	SBold
	SItalic
	SUnderline
	SSlowBlink
	SInverse
	SConceal
	SCrossedOut
	SDoublyUnderline
	SReveal
	SOverline
)

var formatMap = map[LogStyle]int{
	SFDBlack:         30,
	SFDRed:           31,
	SFDGreen:         32,
	SFDYellow:        33,
	SFDBlue:          34,
	SFDMagenta:       35,
	SFDCyan:          36,
	SFDWhite:         37,
	SFBBlack:         90,
	SFBRed:           91,
	SFBGreen:         92,
	SFBYellow:        93,
	SFBBlue:          94,
	SFBMagenta:       95,
	SFBCyan:          96,
	SFBWhite:         97,
	SBDBlack:         40,
	SBDRed:           41,
	SBDGreen:         42,
	SBDYellow:        43,
	SBDBlue:          44,
	SBDMagenta:       45,
	SBDCyan:          46,
	SBDWhite:         47,
	SBBBlack:         100,
	SBBRed:           101,
	SBBGreen:         102,
	SBBYellow:        103,
	SBBBlue:          104,
	SBBMagenta:       105,
	SBBCyan:          106,
	SBBWhite:         107,
	SBold:            1,
	SItalic:          3,
	SUnderline:       4,
	SSlowBlink:       5,
	SInverse:         7,
	SConceal:         8,
	SCrossedOut:      9,
	SDoublyUnderline: 21,
	SReveal:          28,
	SOverline:        53,
}
