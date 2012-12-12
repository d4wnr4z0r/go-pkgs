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
	for _, found := range results {
		fmt.Println(found)
	}
	if len(results) == 1 {
		if results[0] != "./test.d" {
			t.Error("FindDir did not find testdir")
		}
	} else {
		t.Error("length of results is", string(len(results)))
	}
}
