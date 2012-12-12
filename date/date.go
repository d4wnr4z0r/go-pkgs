// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO. 
//
// Package date provides a wrapper around time.Format calls for go's bizarre
// formatting mechanism which is different from the standard mechanism and
// requires reading the time/format.go source to determine.
package date

import "time"

// return mmddyyyy
func MDY(t time.Time) string {
	return t.Format("01022006")
}

// return mm dd yyyy, separated by s
func MDYs(t time.Time, s string) string {
	return t.Format("01"+s+"02"+s+"2006")
}

// return 24-hour time
func HHMM(t time.Time) string {
	return t.Format("1504")
}

