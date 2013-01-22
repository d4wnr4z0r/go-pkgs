// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO. 
//
// Package util provides wrappers around os/io calls for convenience

package util

import (
	"testing"
	"os"
	"../find"
)

func TestIsDir(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	if !IsDir(pwd) {
		t.Error("util.IsDir(pwd) returned false")
	}
}

func TestCopy(t *testing.T) {
	err := Copy("test.file", "test.file.02")
	if err != nil {
		t.Error(err)
	}
}

func TestDirCopy(t *testing.T) {
	err := DirCopy("test.d.1", "test.d.2")
	if err != nil {
		t.Error(err)
	}
	err = os.Chdir("test.d.1")
	if err != nil {
		t.Errorf("Unable to chdir to test.d.1: %s", err)
	}
	test1, err := find.Regex(".", `.`)
	if err != nil {
		t.Error(err)
	}
	err = os.Chdir("../test.d.2")
	if err != nil {
		t.Errorf("Unable to chdir to test.d.2: %s", err)
	}
	test2, err := find.Regex(".", `.`)
	if err != nil {
		t.Error(err)
	}
	for c:= range test1 {
		if test1[c] != test2[c] {
			t.Errorf("test1[%s]: %s, test2[%s]: %s", c, test1[c], c, test2[c])
		}
	}
}
