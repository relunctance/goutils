package fc

import (
	"fmt"
	"testing"
)

func TestUniqid(t *testing.T) {
	v := Uniqid("")
	fmt.Println("len", len(v), "v:", v)
}
