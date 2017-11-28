package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCmdRuncommand(t *testing.T) {
	cmdCode := `ps aux  | grep 'cmd/_test' | grep -v "grep" | grep -v 'go run' | awk '{print $2}'`
	output, err := RunCommand(cmdCode) //get this script run pid
	if err != nil {
		t.Fatalf("Fatal error:%s\n", err)
	}
	if len(output) != 1 {
		t.Fatalf("Fatal error:%s\n", err)
	}

	pid, _ := strconv.Atoi(strings.TrimSpace(output[0]))
	if pid != os.Getpid() { //
		t.Fatalf("pid get error error:exec pid[%d] != os.getpid[%d]\n", pid, os.Getpid())
	}
	filepath := "/tmp/xxx_test_go.txt"
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	f.WriteString("333\n")
	f.WriteString("444\n")
	f.WriteString("555\n")
	defer f.Close()
	if os.IsNotExist(err) {
		t.Fatalf("Fatal error , %s not exists\n", filepath)
	}

	cmdCode = fmt.Sprintf("cat %s | wc -l", filepath)
	Debug = true
	output, err = RunCommand(cmdCode)
	if err != nil {
		t.Fatalf("Fatal error:%s\n", err)
	}

	if output[0] != "3\n" {
		t.Fatalf("run: [cat %s | wc -l ] exec faild\n", filepath)
	}

	if _, err := RunCommand(fmt.Sprintf("rm -f %s", filepath)); err != nil {
		t.Fatalf("rm filepath [%s] error\n", filepath)
	}
}
