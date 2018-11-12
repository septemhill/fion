package fion

import (
	"fmt"
)

const (
	tagTrace = "[TRACE]"
	tagDebug = "[DEBUG]"
	tagInfo  = "[INFO]"
	tagWarn  = "[WARN]"
	tagError = "[ERROR]"
)

var tagFormatMap = map[string]func(format string, a ...interface{}) string{
	tagTrace: NewFormat(SBold, SFBWhite),
	tagDebug: NewFormat(SBold, SFBGreen),
	tagInfo:  NewFormat(SBold, SFBYellow),
	tagWarn:  NewFormat(SBold, SFBBlue),
	tagError: NewFormat(SBold, SFDRed),
}

func composeFormat(tag string, format string, a ...interface{}) string {
	return fmt.Sprintf("%s", tagFormatMap[tag](tag)) + fmt.Sprintf(format, a...)
}

func composeString(tag string, strs ...interface{}) string {
	return fmt.Sprint(tagFormatMap[tag](tag)) + fmt.Sprintln(strs...)
}

func Trace(a ...interface{}) {
	fmt.Print(composeString(tagTrace, a...))
}

func Debug(a ...interface{}) {
	fmt.Print(composeString(tagDebug, a...))
}

func Info(a ...interface{}) {
	fmt.Print(composeString(tagInfo, a...))
}

func Warn(a ...interface{}) {
	fmt.Print(composeString(tagWarn, a...))
}

func Error(a ...interface{}) {
	fmt.Print(composeString(tagError, a...))
}

func Tracef(format string, a ...interface{}) {
	fmt.Println(composeFormat(tagTrace, format, a...))
}

func Debugf(format string, a ...interface{}) {
	fmt.Println(composeFormat(tagDebug, format, a...))
}

func Infof(format string, a ...interface{}) {
	fmt.Println(composeFormat(tagInfo, format, a...))
}

func Warnf(format string, a ...interface{}) {
	fmt.Println(composeFormat(tagWarn, format, a...))
}

func Errorf(format string, a ...interface{}) {
	fmt.Println(composeFormat(tagError, format, a...))
}
