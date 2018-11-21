package fc

import (
	"testing"
)

func TestRangeByte(t *testing.T) {

	if string(RangeByte('a', 'm', 1)) != "abcdefghijklm" {
		t.Fatalf("should be equal")
	}
	if string(RangeByte('a', 'z', 2)) != "acegikmoqsuwy" {
		t.Fatalf("should be equal")
	}

}

func TestRangeInt(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	if !checkIntSlice(a, RangeInt(0, 10, 1)) {
		t.Fatalf("should be equal %#v\n", a)
	}

	if !checkIntSlice(b, RangeInt(0, 20, 2)) {
		t.Fatalf("should be equal %#v\n", b)
	}
}

func checkIntSlice(a, b []int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
