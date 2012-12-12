// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO.                                    
// Package slog provides an interface to syslog which allows setting the syslog
// facillity.
package slog

/*
#include <stdlib.h>
#include <syslog.h>
void local_syslog (int p, const char *f) {
	syslog (p, f);
}
*/
import "C"
import (
	"unsafe"
	"errors"
)

// map of facilities to their C counterparts
var facilities = map[string]int {
	"kern":		C.LOG_KERN,
	"user":		C.LOG_USER,
	"mail":		C.LOG_MAIL,
	"daemon":	C.LOG_DAEMON,
	"auth":		C.LOG_AUTH,
	"syslog":	C.LOG_SYSLOG,
	"lpr":		C.LOG_LPR,
	"news":		C.LOG_NEWS,
	"uucp":		C.LOG_UUCP,
	"cron":		C.LOG_CRON,
	"authpriv":	C.LOG_AUTHPRIV,
	"ftp":		C.LOG_FTP,
	"local0":	C.LOG_LOCAL0,
	"local1":	C.LOG_LOCAL1,
	"local3":	C.LOG_LOCAL3,
	"local4":	C.LOG_LOCAL4,
	"local5":	C.LOG_LOCAL5,
	"local6":	C.LOG_LOCAL6,
	"local7":	C.LOG_LOCAL7,
}

// map of priorities to their C counterparts
var priorities = map[string]int {
	"emerg":	C.LOG_EMERG,
	"alert":	C.LOG_ALERT,
	"crit":		C.LOG_CRIT,
	"err":		C.LOG_ERR,
	"warn":		C.LOG_WARNING,
	"notice":	C.LOG_NOTICE,
	"info":		C.LOG_INFO,
	"debug":	C.LOG_DEBUG,
}

var facility, priority C.int
var name string

// function to set facility to log to
func SetFac(s string) error {
	if s == "" {
		return errors.New("received an empty string as a facility")
	}
	facility = C.int(facilities[s])
	return nil
}

// function to set the logging-as name
func SetName(s string) error {
	if s == "" {
		return errors.New("received an emptry string as a name")
	}
	name = s
	return nil
}

// send sends a given string to syslog 
func send(m string) error {
	if m == "" {
		return errors.New("func send received an emptry string to log.")
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.openlog(c_name, 0, facility)

	c_str := C.CString(m)
	defer C.free(unsafe.Pointer(c_str))
	C.local_syslog(priority, c_str)

	return nil
}

func sendPriority(p string, m string) (err error) {
	priority = C.int(priorities[p])
	err = send(m)
	return
}

// send an emergency to syslog
func Emerg(m string) (err error) {
	err = sendPriority("emerge", m)
	return
}

// send an alert to syslog
func Alert(m string) (err error) {
	err = sendPriority("alert", m)
	return
}

// send a critical to syslog
func Crit(m string) (err error) {
	err = sendPriority("crit", m)
	return
}

// send an error to syslog
func Err(m string) (err error) {
	err = sendPriority("err", m)
	return
}

// send a warning to syslog
func Warn(m string) (err error) {
	err = sendPriority("warn", m)
	return
}

// send a notice to syslog
func Notice(m string) (err error) {
	err = sendPriority("notice", m)
	return
}

// send an info to syslog
func Info(m string) (err error) {
	err = sendPriority("info", m)
	return
}

// send a debug to syslog
func Debug(m string) (err error) {
	err = sendPriority("debug", m)
	return
}
