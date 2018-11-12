package fion

//Black make format string foreground to be black color
func Black(format string, a ...interface{}) string {
	return generateFormat(SFDBlack, format, a...)
}

//Red make format string foreground to be red color
func Red(format string, a ...interface{}) string {
	return generateFormat(SFDRed, format, a...)
}

//Green make format string foreground to be green color
func Green(format string, a ...interface{}) string {
	return generateFormat(SFDGreen, format, a...)
}

//Yellow make format string foreground to be yellow color
func Yellow(format string, a ...interface{}) string {
	return generateFormat(SFDYellow, format, a...)
}

//Blue make format string foreground to be blue color
func Blue(format string, a ...interface{}) string {
	return generateFormat(SFDBlue, format, a...)
}

//Magenta make format string foreground to be magenta color
func Magenta(format string, a ...interface{}) string {
	return generateFormat(SFDMagenta, format, a...)
}

//Cyan make format string foreground to be cyan color
func Cyan(format string, a ...interface{}) string {
	return generateFormat(SFDCyan, format, a...)
}

//White make format string foreground to be white color
func White(format string, a ...interface{}) string {
	return generateFormat(SFDWhite, format, a...)
}

//BBlack make format string background to be black color
func BBlack(format string, a ...interface{}) string {
	return generateFormat(SBDBlack, format, a...)
}

//BRed make format string background to be red color
func BRed(format string, a ...interface{}) string {
	return generateFormat(SBDRed, format, a...)
}

//BGreen make format string background to be green color
func BGreen(format string, a ...interface{}) string {
	return generateFormat(SBDGreen, format, a...)
}

//BYellow make format string background to be yellow color
func BYellow(format string, a ...interface{}) string {
	return generateFormat(SBDYellow, format, a...)
}

//BBlue make format string background to be blue color
func BBlue(format string, a ...interface{}) string {
	return generateFormat(SBDBlue, format, a...)
}

//BMagenta make format string background to be magenta color
func BMagenta(format string, a ...interface{}) string {
	return generateFormat(SBDMagenta, format, a...)
}

//BCyan make format string background to be cyan color
func BCyan(format string, a ...interface{}) string {
	return generateFormat(SBDCyan, format, a...)
}

//BWhite make format string background to be white color
func BWhite(format string, a ...interface{}) string {
	return generateFormat(SBDWhite, format, a...)
}
