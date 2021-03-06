package fion

import (
	"fmt"
	"os"
	"regexp"
	"syscall"
	"time"
)

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

func logToFile(msg string) {
	now := time.Now()
	re := regexp.MustCompile("\x1b[[0-9;]*m")
	text := re.ReplaceAllLiteralString(msg, "")
	file, _ := os.OpenFile(fionConf.errorLog, syscall.O_CREAT|syscall.O_APPEND|syscall.O_RDWR, 0664)
	file.Write([]byte(now.String() + text))
	file.Close()
}

//SetTagLevel rewrite tag level
func SetTagLevel(lv int) {
	fionConf.setTagLevel(lv)
}

//EnableExternalLog enable output stream redirect to log file
func EnableExternalLog() {
	fionConf.enableLog()
}

//DisableExternalLog disable output stream redirect to log file
func DisableExternalLog() {
	fionConf.disableLog()
}

//NewTag customize your own tag funcion
func NewTag(tagName string, curTagLevel int, f ...LogStyle) (Print, Printf) {
	return func(a ...interface{}) {
			if curTagLevel >= fionConf.tagLevel {
				fmt.Print(composeString(NewFormat(f...)(tagName), a...))
			}
			if fionConf.externalLog {
				logToFile(composeString(NewFormat(f...)(tagName), a...))
			}
		}, func(format string, a ...interface{}) {
			if curTagLevel >= fionConf.tagLevel {
				fmt.Println(composeFormat(NewFormat(f...)(tagName), format, a...))
			}
			if fionConf.externalLog {
				logToFile(composeFormat(NewFormat(f...)(tagName), format, a...) + "\n")
			}
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
