package fc

import (
	"fmt"
	"github.com/fatih/color"
)

// c := color.New(color.FgCyan).Add(color.Underline)
// c := color.New(color.FgCyan, color.Bold)
// https://github.com/fatih/color
func ColorDump(c *color.Color, vals ...interface{}) {
	if c == nil {
		Dump(vals...)
		return
	}
	for _, v := range vals {
		c.Println(JsonDump(v))
	}
}

// 类似PHP  var_dump
func Dump(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(JsonDump(v))
	}
}
