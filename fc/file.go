package fc

import (
    "os"
    "path/filepath"
)

func IsExist(filename string) bool {
    if _, err := os.Stat(filename); err == nil {
        return true
    }
    return false
}

func FileSize(filename string) int64 {
    var size int64
    filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
        size = f.Size()
        return nil
    })
    return size
}
