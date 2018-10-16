package fc

import (
	"github.com/axgle/mahonia"
)

func GbkToUtf8(v string) string {
	dec := mahonia.NewDecoder("gbk") //decode
	ret := dec.ConvertString(v)
	return ret

}

func Utf8ToGbk(v string) string {
	enc := mahonia.NewEncoder("gbk") //encode
	ret := enc.ConvertString(v)
	return ret
}
