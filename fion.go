package fion

import "fmt"

func buildFormatting(f ...FormatNumber) string {
	format := ""

	for _, value := range f {
		//fmt.Println("RRR", key, value)
		format += formatMap[value]
	}

	return format
}

func New(f ...FormatNumber) func(format string, a ...interface{}) string {
	//str := ";1;3;5;31"
	customizeFormat := buildFormatting(f...)
	fmt.Println("[FORMAT]", customizeFormat)

	return func(format string, a ...interface{}) string {
		if len(a) > 0 {
			return fmt.Sprintf(fmt.Sprintf("\x1b[%sm%s\x1b[0m", customizeFormat, format), a...)
		}
		return fmt.Sprintf("\x1b[%sm%s\x1b[0m", customizeFormat, format)
	}
}
