package fion

import "fmt"

type styleNumber uint8

const (
	styleItalic          = styleNumber(3)
	styleUnderline       = styleNumber(4)
	styleSlowBlink       = styleNumber(5)
	styleInverse         = styleNumber(7)
	styleConceal         = styleNumber(8)
	styleCrossedOut      = styleNumber(9)
	styleDoublyUnderline = styleNumber(21)
	styleReveal          = styleNumber(28)
	styleOverline        = styleNumber(53)
)

func generalFormat2(style styleNumber, format string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(fmt.Sprintf("\x1b[%d;0m%s\x1b[0m", style, format), a...)
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", style, format)
}

func Italic(format string, a ...interface{}) string {
	return generalFormat2(styleItalic, format, a...)
}

func Underline(format string, a ...interface{}) string {
	return generalFormat2(styleUnderline, format, a...)
}

func SlowBlink(format string, a ...interface{}) string {
	return generalFormat2(styleSlowBlink, format, a...)
}

func Inverse(format string, a ...interface{}) string {
	return generalFormat2(styleInverse, format, a...)
}

func Conceal(format string, a ...interface{}) string {
	return generalFormat2(styleConceal, format, a...)
}

func CrossedOut(format string, a ...interface{}) string {
	return generalFormat2(styleCrossedOut, format, a...)
}

func DoublyUnderline(format string, a ...interface{}) string {
	return generalFormat2(styleDoublyUnderline, format, a...)
}

func Reveal(format string, a ...interface{}) string {
	return generalFormat2(styleReveal, format, a...)
}

func Overline(format string, a ...interface{}) string {
	return generalFormat2(styleOverline, format, a...)
}
