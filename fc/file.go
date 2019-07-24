package fc

import (
	"bytes"
	"io/ioutil"
	"os"
	"syscall"
)

// check filepath is exists
// if exists return true else return false
func IsExist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

// check file or dir is writeable
func IsWriteable(name string) bool {
	err := syscall.Access(name, syscall.O_RDWR)
	if err == nil {
		return true
	}
	return false
}

// check file or dir is readable
func IsReadable(name string) bool {
	err := syscall.Access(name, syscall.O_RDONLY)
	if err == nil {
		return true
	}
	return false
}

// get file size by filepath
func FileSize(filename string) int64 {
	if info, err := os.Stat(filename); err == nil {
		return info.Size()
	}

	return 0
}

func FileGet(filename string) ([][]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.Split(content, []byte("\n")), nil
}

func FileStrings(filename string) ([]string, error) {
	lines, err := FileGet(filename)
	if err != nil {
		return nil, err
	}
	data := make([]string, 0, len(lines))
	for _, line := range lines {
		data = append(data, string(line))
	}
	return data, nil
}
