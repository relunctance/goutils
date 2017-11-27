package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"unicode"
)

//是否开启debug
var Debug bool = false

//允许管道出现的最大次数
var SPLIT_NUM int = 20

//运行命令, 支持管道
//示例1: RunCommand("ls /home/")
//示例2: RunCommand("ps auxwww | grep init | grep -v grep")
func RunCommand(cmd string) ([]string, error) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return nil, errors.New("cmd code is empty")
	}
	var cmdSlice []string
	if strings.Index(cmd, "|") != 0 {
		cmdSlice = strings.SplitN(cmd, "|", SPLIT_NUM)
	} else {
		cmdSlice = append(cmdSlice, cmd)
	}

	cmds := make([]*exec.Cmd, len(cmdSlice))
	for i, cmdCode := range cmdSlice {
		cmdCode = strings.TrimSpace(cmdCode)
		cmdItem := CmdFields(cmdCode)

		if Debug {
			fmt.Printf("run cmdFields: %#v\n", cmdItem)
		}
		if len(cmdItem) == 1 {
			cmds[i] = exec.Command(cmdItem[0])
		} else {
			cmds[i] = exec.Command(cmdItem[0], cmdItem[1:]...)
		}

	}
	return runCmds(cmds)
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {

	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("The cmd slice is invalid!")
	}
	first := true
	var output []byte
	var err error
	for _, cmd := range cmds {
		if Debug {
			fmt.Printf("Run command: %v\n", getCmdPlaintext(cmd)) //可以查看命令的绝对路径
		}
		if !first {
			var stdinBuf bytes.Buffer
			stdinBuf.Write(output) //非首次情况下, 写入output
			cmd.Stdin = &stdinBuf
		}
		var stdoutBuf bytes.Buffer
		cmd.Stdout = &stdoutBuf
		if err = cmd.Start(); err != nil {
			return nil, getError(err, cmd)
		}
		if err = cmd.Wait(); err != nil {
			return nil, getError(err, cmd)
		}
		output = stdoutBuf.Bytes() //无论首次,还是多个管道,都写入output
		//fmt.Printf("Output:\n%s\n", string(output))
		if first {
			first = false
		}
	}
	var lines []string
	var outputBuf bytes.Buffer
	outputBuf.Write(output)
	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, getError(err, nil)
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func getCmdPlaintext(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path)
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String()
}

func getError(err error, cmd *exec.Cmd, extraInfo ...string) error {
	var errMsg string
	if cmd != nil {
		errMsg = fmt.Sprintf("%s  [%s %v]", err, (*cmd).Path, (*cmd).Args)
	} else {
		errMsg = fmt.Sprintf("%s", err)
	}
	if len(extraInfo) > 0 {
		errMsg = fmt.Sprintf("%s (%v)", errMsg, extraInfo)
	}
	return errors.New(errMsg)
}

type codeParseSt struct {
	value    string
	isQuotes bool //true 不需要再使用space切割 , false需要
}

//支持cmd命令中单引号 , 双引号拆分
func CmdFields(str string) []string {
	return parseCodeStr(str)
}

func parseCodeStr(str string) []string {
	arr := splitByQuotes(str)
	codeSlice := getCodeParseStruct(str, arr)
	codeArray := make([]string, 0)
	for _, codeSt := range codeSlice {
		if (*codeSt).isQuotes {
			codeArray = append(codeArray, (*codeSt).value)
		} else {
			tmpslice := strings.Fields((*codeSt).value)
			codeArray = append(codeArray, tmpslice...)
		}
	}
	return codeArray
}

func splitByQuotes(str string) []string {
	arr := strings.FieldsFunc(str, func(c rune) bool {
		if uint32(c) <= unicode.MaxLatin1 {
			return checkIsQuotes(c)
		}
		return false
	})
	return arr
}

func getCodeParseStruct(str string, arr []string) []*codeParseSt {
	codeSlice := make([]*codeParseSt, 0, len(arr))
	for _, val := range arr {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		i := strings.Index(str, val)
		if i == 0 {
			codeSlice = append(codeSlice, &codeParseSt{val, false})
			continue
		}
		i--
		isQuotes := checkIsQuotes(rune(str[i]))
		codeSlice = append(codeSlice, &codeParseSt{val, isQuotes})
	}
	return codeSlice
}

func checkIsQuotes(c rune) bool {
	switch c {
	case '"', '\'':
		return true
	}
	return false
}
