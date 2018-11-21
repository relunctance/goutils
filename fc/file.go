package fc

import (
	"os"
)

// check filepath is exists
// if exists return true else return false
func IsExist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

// get file size by filepath
func FileSize(filename string) int64 {
	if info, err := os.Stat(filename); err == nil {
		return info.Size()
	}

	return 0
}
