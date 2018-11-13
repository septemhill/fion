package fion

import (
	"fmt"
)

var Trace Print
var Debug Print
var Info Print
var Warn Print
var Error Print

var Tracef Printf
var Debugf Printf
var Infof Printf
var Warnf Printf
var Errorf Printf

func composeFormat(tag string, format string, a ...interface{}) string {
	return fmt.Sprintf("%s", tag) + fmt.Sprintf(format, a...)
}

func composeString(tag string, strs ...interface{}) string {
	return fmt.Sprint(tag) + fmt.Sprintln(strs...)
}

func init() {
	t1, t2 := NewTag("[TRACE]", SBold, SFBWhite, SInverse)
	d1, d2 := NewTag("[DEBUG]", SBold, SFBGreen, SInverse)
	i1, i2 := NewTag("[INFO]", SBold, SFBYellow, SInverse)
	w1, w2 := NewTag("[WARN]", SBold, SFBBlue, SInverse)
	e1, e2 := NewTag("[ERROR]", SBold, SFDRed, SInverse)

	Trace, Debug, Info, Warn, Error = t1, d1, i1, w1, e1
	Tracef, Debugf, Infof, Warnf, Errorf = t2, d2, i2, w2, e2
}
