package fc

import (
	"strings"
	"testing"
)

func TestDebugTrace(t *testing.T) {
	ts := DebugTrace()
	newts := getTraceInfo(ts)
	result := []string{
		"debug.go:debug.DebugTrace()",
		"debug_test.go:debug.TestDebugTrace()",
		"testing.go:testing.tRunner()",
		"asm_amd64.s:runtime.goexit()",
	}

	if !CheckStringSliceEqual(newts, result) {
		t.Errorf("expcect true newts:\n %v\n === \nresult:\n%v\n", newts, result)
	}
}

func getTraceInfo(arr []string) (ret []string) {

	for _, val := range arr {
		strings.Index(val, "/")
		pos := strings.LastIndex(val, "/")
		if pos == -1 {
			continue
		}
		v := strings.TrimSpace(val[pos+1:])
		tmparr := strings.Split(v, ":")
		ret = append(ret, tmparr[0]+":"+tmparr[2])
	}
	return
}
