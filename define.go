package fion

type FormatNumber int

//FD: Foreground Default
//FB: Foreground Bright
//BD: Background Default
//BB: Background Bright
const (
	FDBlack FormatNumber = iota
	FDRed
	FDGreen
	FDYellow
	FDBlue
	FDMagenta
	FDCyan
	FDWhite
	FBBlack
	FBRed
	FBGreen
	FBYellow
	FBBlue
	FBMagenta
	FBCyan
	FBWhite
	BDBlack
	BDRed
	BDGreen
	BDYellow
	BDBlue
	BDMagenta
	BDCyan
	BDWhite
	BBBlack
	BBRed
	BBGreen
	BBYellow
	BBBlue
	BBMagenta
	BBCyan
	BBWhite
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

var formatMap = map[FormatNumber]string{
	FDBlack:          ";30",
	FDRed:            ";31",
	FDGreen:          ";32",
	FDYellow:         ";33",
	FDBlue:           ";34",
	FDMagenta:        ";35",
	FDCyan:           ";36",
	FDWhite:          ";37",
	FBBlack:          ";90",
	FBRed:            ";91",
	FBGreen:          ";92",
	FBYellow:         ";93",
	FBBlue:           ";94",
	FBMagenta:        ";95",
	FBCyan:           ";96",
	FBWhite:          ";97",
	BDBlack:          ";40",
	BDRed:            ";41",
	BDGreen:          ";42",
	BDYellow:         ";43",
	BDBlue:           ";44",
	BDMagenta:        ";45",
	BDCyan:           ";46",
	BDWhite:          ";47",
	BBBlack:          ";100",
	BBRed:            ";101",
	BBGreen:          ";102",
	BBYellow:         ";103",
	BBBlue:           ";104",
	BBMagenta:        ";105",
	BBCyan:           ";106",
	BBWhite:          ";107",
	SBold:            ";1",
	SItalic:          ";3",
	SUnderline:       ";4",
	SSlowBlink:       ";5",
	SInverse:         ";7",
	SConceal:         ";8",
	SCrossedOut:      ";9",
	SDoublyUnderline: ";21",
	SReveal:          ";28",
	SOverline:        ";53",
}
