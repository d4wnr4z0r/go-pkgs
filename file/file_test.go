// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO. 
//
// Package file provides wrappers around os/io calls for convenience

package file

import (
	"testing"
	"os"
)

func TestIsDir(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	if !IsDir(pwd) {
		t.Error("file.IsDir(pwd) returned false")
	}
}

func TestCopy(t *testing.T) {
	err := Copy("test.file", "test.file.02")
	if err != nil {
		t.Error(err)
	}
}
