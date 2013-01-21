// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO.
package find

import (
	"testing"
)

func TestRegex01(t *testing.T) {
	results, err := Regex(".", `\.go$`)
	if err != nil {
		t.Errorf("Find returned error: %s", err)
	}
	if len(results) == 2 {
		if results[0] != "./find.go" {
			t.Error("Find found "+ results[0] +" instead of find_test.go")
		}
		if results[1] != "./find_test.go" {
			t.Error("Find found "+ results[1] +" instead of find.go")
		}
	} else {
		t.Errorf("length of results is %d", len(results))
	}
}

// invalid input tests
func TestRegex02(t *testing.T) {
	_, err := Regex("/omg/wtf/bbq", `*\.*`)
	if err == nil {
		t.Error("did not receive err for invalid regex")
	}
	_, err = Regex("/omg/wtf/bbq", `\.*`)
	if err == nil {
		t.Error("did not receive err for invalid directory")
	}
}

func TestDirRegex01(t *testing.T) {
	results, err := DirRegex(".", `\.d$`)
	if err != nil {
		t.Errorf("DirRegex(\".\", `\\.d$`) produced error %s", err)
	}
	if len(results) == 1 {
		if results[0] != "./test.d" {
			t.Error("DirRegex found "+ results[0] +" instead of test.d")
		}
	} else {
		t.Errorf("length of results is %d", len(results))
	}
}

