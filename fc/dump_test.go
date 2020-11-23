package fc

import (
	"testing"

	"github.com/fatih/color"
)

func TestDump(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	Dump(
		[]string{"a", "b", "c"},
		"abc",
		1,
		1.68,
		[]int{1, 2, 3, 4, 5},
		[]int32{1, 2, 3, 4, 5},
		[]User{
			User{"hello", 12},
			User{"world", 13},
		},
		[]*User{
			&User{"hello", 14},
			&User{"world", 15},
		},
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	)
}

func TestColorDump(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	ColorDump(
		color.New(color.FgCyan, color.Bold),
		[]string{"a", "b", "c"},
		"abc",
		1,
		1.68,
		[]int{1, 2, 3, 4, 5},
		[]int32{1, 2, 3, 4, 5},
		[]User{
			User{"hello", 12},
			User{"world", 13},
		},
		[]*User{
			&User{"hello", 14},
			&User{"world", 15},
		},
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	)
}
