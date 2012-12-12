// copyright 2012 Jacob Pipkin                                                  
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO.
//
// Package find provides (a) mechanism(s) to search for files/folders under
// a directory tree.
package find

import (
	"testing"
	"fmt"
)

func TestRegex01(t *testing.T) {
	results, err := Regex(".", `\.go$`)
	if err != nil {
		t.Errorf("Find returned error: %s", err)
	}
	if len(results) == 2 {
		if results[0] != "./find_test.go" {
			t.Error("Find did not find find_test.go")
		}
		if results[1] != "./find.go" {
			t.Error("Find did not find find.go")
		}
	} else {
		for _, found := range results {
			fmt.Println(found)
		}
		t.Errorf("length of results is %s", len(results))
	}
}

// invalid input tests
func TestRegex02(t *testing.T) {
	_, err := Regex("/omg/wtf/bbq", `*\.*`)
	if err == nil {
		t.Error("did not receive err for invalid regex")
	} else {
		fmt.Println(err)
	}
	_, err = Regex("/omg/wtf/bbq", `\.*`)
	if err == nil {
		t.Error("did not receive err for invalid directory")
	} else {
		fmt.Println(err)
	}
}

func TestDirRegex01(t *testing.T) {
	results, err := DirRegex(".", `dir$`)
	if err != nil {
		t.Errorf("FindDir returned error: %s", err)
	}
	if len(results) == 1 {
		if results[0] != "./testdir" {
			t.Error("FindDir did not find testdir")
		}
	} else {
		for _, found := range results {
			fmt.Println(found)
		}
		t.Error("length of results is", string(len(results)))
	}
}
