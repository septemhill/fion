package fion

import (
	"fmt"
)

const (
	prefixTrace = "[Trace]"
	prefixDebug = "[Debug]"
	prefixInfo  = "[Info]"
	prefixWarn  = "[WARN]"
	prefixError = "[ERROR]"
)

var prefixMap = map[string]func(format string, a ...interface{}) string{
	prefixTrace: White,
	prefixDebug: Green,
	prefixInfo:  Yellow,
	prefixWarn:  Blue,
	prefixError: Red,
}

func composeFormat(prefix string, format string, a ...interface{}) string {
	return fmt.Sprintf("%s", prefixMap[prefix](prefix)) + fmt.Sprintf(format, a...)
}

func composeString(prefix string, strs ...interface{}) string {
	return fmt.Sprint(prefixMap[prefix](prefix)) + fmt.Sprintln(strs...)
}

func Trace(a ...interface{}) {
	fmt.Print(composeString(prefixTrace, a...))
}

func Debug(a ...interface{}) {
	fmt.Print(composeString(prefixDebug, a...))
}

func Info(a ...interface{}) {
	fmt.Print(composeString(prefixInfo, a...))
}

func Warn(a ...interface{}) {
	fmt.Print(composeString(prefixWarn, a...))
}

func Error(a ...interface{}) {
	fmt.Print(composeString(prefixError, a...))
}

func Tracef(format string, a ...interface{}) {
	fmt.Println(composeFormat(prefixTrace, format, a...))
}

func Debugf(format string, a ...interface{}) {
	fmt.Println(composeFormat(prefixDebug, format, a...))
}

func Infof(format string, a ...interface{}) {
	fmt.Println(composeFormat(prefixInfo, format, a...))
}

func Warnf(format string, a ...interface{}) {
	fmt.Println(composeFormat(prefixWarn, format, a...))
}

func Errorf(format string, a ...interface{}) {
	fmt.Println(composeFormat(prefixError, format, a...))
}
