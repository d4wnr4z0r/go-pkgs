// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
package slog

import "testing"

func TestSetFac01(t *testing.T) {
	in, out := "local7", 184
	SetFac(in)
	result := int(facility)
	if result != out {
		t.Errorf("SetFac(%v), facility == %v, expected %v", in, result, out)
	}
}

func TestSetName01(t *testing.T) {
	in, out := "test_name", "test_name"
	SetName(in)
	result := name
	if result != out {
		t.Errorf("SetName(%v), name == %v, expected %v", in, result, out)
	}
}

// test sending any empty string to send()
func Testsend01(t *testing.T) {
	err := send("")
	if err == nil {
		t.Error("send allows an empty string")
	}
}
