package chapter42

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	var a, b string
	var i int

	a, b = "", ""
	i = 0
	cmp := strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but compare(%q, %q) = %v", i, a, b, cmp)
	}

	a, b = "a", ""
	i = 1
	cmp = strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but compare(%q, %q) = %v", i, a, b, cmp)
	}

	a, b = "", "a"
	i = -1
	cmp = strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but compare(%q, %q) = %v", i, a, b, cmp)
	}
}
