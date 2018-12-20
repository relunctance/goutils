package fc

import (
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
