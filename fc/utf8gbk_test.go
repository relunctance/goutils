package fc

import (
	"testing"
)

func TestUtf8ToGbk(t *testing.T) {

	str := "你好，世界！"
	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	v := Utf8ToGbk(str)
	if v != string(testBytes) {
		t.Fatalf("[%v]should be =='%v'", []byte(v), testBytes)
	}
}

func TestGbkToUtf8(t *testing.T) {
	str := "你好，世界！"
	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	v := GbkToUtf8(string(testBytes))
	if v != str {
		t.Fatalf("[%s]should be =='%s'", v, str)
	}

}
