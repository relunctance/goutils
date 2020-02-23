package fc

import (
	"testing"
)

func TestCoverUtf8(t *testing.T) {
	arr := map[string]string{
		"\u7ee0\uff04\u608a\u9a9e\u51b2\u5f74\u9427\u5a5a\u6ab0":         "管理平台登陆",
		"\u63a8\u8350\u57281024x768\u5206\u8fa8\u7387\u4e0b\u64cd\u4f5c": "推荐在1024x768分辨率下操作",
	}
	for v, check := range arr {
		v = UnicodeToString(v)
		if IsCanUtf8ToGbk(v) {
			v = Utf8ToGbk(v)
		}
		if v != check {
			t.Fatalf("vv:[%s] should be == [%s]", v, check)
		}
	}
}

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
