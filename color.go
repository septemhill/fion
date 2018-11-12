package fion

import "fmt"

type colorNumber uint8

const (
	backgroundColorBlack = colorNumber(iota + 100)
	backgroundColorRed
	backgroundColorGreen
	backgroundColorYellow
	backgroundColorBlue
	backgroundColorMagenta
	backgroundColorCyan
	backgroundColorWhite
)

const (
	foregroundColorBlack = colorNumber(iota + 90)
	foregroundColorRed
	foregroundColorGreen
	foregroundColorYellow
	foregroundColorBlue
	foregroundColorMagenta
	foregroundColorCyan
	foregroundColorWhite
)

func generalFormat(color colorNumber, format string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, format), a...)
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, format)
}

//Black make format string foreground to be black color
func Black(format string, a ...interface{}) string {
	return generalFormat(foregroundColorBlack, format, a...)
}

//Red make format string foreground to be red color
func Red(format string, a ...interface{}) string {
	return generalFormat(foregroundColorRed, format, a...)
}

//Green make format string foreground to be green color
func Green(format string, a ...interface{}) string {
	return generalFormat(foregroundColorGreen, format, a...)
}

//Yellow make format string foreground to be yellow color
func Yellow(format string, a ...interface{}) string {
	return generalFormat(foregroundColorYellow, format, a...)
}

//Blue make format string foreground to be blue color
func Blue(format string, a ...interface{}) string {
	return generalFormat(foregroundColorBlue, format, a...)
}

//Magenta make format string foreground to be magenta color
func Magenta(format string, a ...interface{}) string {
	return generalFormat(foregroundColorMagenta, format, a...)
}

//Cyan make format string foreground to be cyan color
func Cyan(format string, a ...interface{}) string {
	return generalFormat(foregroundColorCyan, format, a...)
}

//White make format string foreground to be white color
func White(format string, a ...interface{}) string {
	return generalFormat(foregroundColorWhite, format, a...)
}

//BBlack make format string background to be black color
func BBlack(format string, a ...interface{}) string {
	return generalFormat(backgroundColorBlack, format, a...)
}

//BRed make format string background to be red color
func BRed(format string, a ...interface{}) string {
	return generalFormat(backgroundColorRed, format, a...)
}

//BGreen make format string background to be green color
func BGreen(format string, a ...interface{}) string {
	return generalFormat(backgroundColorGreen, format, a...)
}

//BYellow make format string background to be yellow color
func BYellow(format string, a ...interface{}) string {
	return generalFormat(backgroundColorYellow, format, a...)
}

//BBlue make format string background to be blue color
func BBlue(format string, a ...interface{}) string {
	return generalFormat(backgroundColorBlue, format, a...)
}

//BMagenta make format string background to be magenta color
func BMagenta(format string, a ...interface{}) string {
	return generalFormat(backgroundColorMagenta, format, a...)
}

//BCyan make format string background to be cyan color
func BCyan(format string, a ...interface{}) string {
	return generalFormat(backgroundColorCyan, format, a...)
}

//BWhite make format string background to be white color
func BWhite(format string, a ...interface{}) string {
	return generalFormat(backgroundColorWhite, format, a...)
}
