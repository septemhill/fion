package fion

import "fmt"

func generateFormat(style LogStyle, format string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(fmt.Sprintf("\x1b[%dm%s\x1b[0m", formatMap[style], format), a...)
	}

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", formatMap[style], format)
}

func buildFormatting(f ...LogStyle) string {
	format := ""

	for _, value := range f {
		format += fmt.Sprintf(";%d", formatMap[value])
	}

	return format
}

//NewTag customize your own tag funcion
func NewTag(tagName string, f ...LogStyle) (Print, Printf) {
	return func(a ...interface{}) {
			fmt.Print(composeString(NewFormat(f...)(tagName), a...))
		}, func(format string, a ...interface{}) {
			fmt.Println(composeFormat(NewFormat(f...)(tagName), format, a...))
		}
}

//NewFormat customize your own format function
func NewFormat(f ...LogStyle) func(format string, a ...interface{}) string {
	customizeFormat := buildFormatting(f...)

	return func(format string, a ...interface{}) string {
		if len(a) > 0 {
			return fmt.Sprintf(fmt.Sprintf("\x1b[%sm%s\x1b[0m", customizeFormat, format), a...)
		}
		return fmt.Sprintf("\x1b[%sm%s\x1b[0m", customizeFormat, format)
	}
}
