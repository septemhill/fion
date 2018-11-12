package fion

//Bold make string format to be bold
func Bold(format string, a ...interface{}) string {
	return generateFormat(SBold, format, a...)
}

//Italic make string format to be italic
func Italic(format string, a ...interface{}) string {
	return generateFormat(SItalic, format, a...)
}

//Underline make string format to be underline
func Underline(format string, a ...interface{}) string {
	return generateFormat(SUnderline, format, a...)
}

//SlowBlink make string format to be slow blink
func SlowBlink(format string, a ...interface{}) string {
	return generateFormat(SSlowBlink, format, a...)
}

//Inverse make string format to be inverse
func Inverse(format string, a ...interface{}) string {
	return generateFormat(SInverse, format, a...)
}

//Conceal make string format to be conceal
func Conceal(format string, a ...interface{}) string {
	return generateFormat(SConceal, format, a...)
}

//CrossedOut make string format to be crossed out
func CrossedOut(format string, a ...interface{}) string {
	return generateFormat(SCrossedOut, format, a...)
}

//DoublyUnderline make string format to be doubly underline
func DoublyUnderline(format string, a ...interface{}) string {
	return generateFormat(SDoublyUnderline, format, a...)
}

//Reveal make string format to be reveal
func Reveal(format string, a ...interface{}) string {
	return generateFormat(SReveal, format, a...)
}

//Overline make string format to be overline
func Overline(format string, a ...interface{}) string {
	return generateFormat(SOverline, format, a...)
}
