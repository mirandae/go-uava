package check

import "testing"

var notEmptyString = []struct {
	in  string
	exp error
}{
	{"", nil},
	{"test", nil},
}

func TestNotEmptyString(t *testing.T) {}
