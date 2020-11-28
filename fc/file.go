package fc

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
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

// SaveFile data save to file
func SaveFile(path string, content []byte, bak bool) error {
	tmp := path + ".tmp"
	err := AttachFileDirectory(tmp)
	if err != nil {
		return err
	}
	if bak && IsExist(path) {
		// bak
		_ = CopyFileContents(path, path+".bak")
	}
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	_ = f.Close()
	return os.Rename(tmp, path)
}

func AttachFileDirectory(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		_ = in.Close()
	}()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
