package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

var input string
var output string
var cover bool
var checksize bool

func main() {
	flag.StringVar(&input, "input", "", "the copy input dir , the type should be dir")
	flag.StringVar(&output, "output", "", "the copy input dir, the type should be dir")
	flag.BoolVar(&cover, "cover", false, "is cover exists file")
	flag.BoolVar(&checksize, "check", true, "is check size when use -cover=true")
	flag.Parse()
	fmt.Println("---------------------------------\n")
	fmt.Println("input:", input)
	fmt.Println("output:", output)
	fmt.Println("cover:", cover)
	fmt.Println("check:", checksize)
	fmt.Println("\n---------------------------------\n")
	input = strings.TrimRight(input, "/")
	output = strings.TrimRight(output, "/")
	if input == "" || output == "" {
		flag.Usage()
		return
	}
	if err := checkPath(input); err != nil {
		panic(err)
	}

	if err := checkPath(output); err != nil {
		panic(err)
	}

	names := fileNames(input)
	l := len(names)
	if l == 0 {
		log.Printf("input dir: [%s] is a empty dir\n", input)
		return
	}

	ch := make(chan string, l)
	for _, name := range names {
		name = strings.TrimRight(name, "\n")

		go func(filename string) {
			err := copyFile(filename, input, output, ch)
			if err != nil {
				panic(err)
			}
		}(name)

	}

	for i := 0; i < l; i++ {
		fmt.Printf("[%s] is ok\n", <-ch)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func copyFile(filename string, input, output string, ch chan string) error {
	if !IsWriteable(output) {
		return fmt.Errorf("[%s] is not writeable \n", output)
	}
	src := input + "/" + filename
	if !IsExist(src) {
		return fmt.Errorf("not exists [%s]\n", src)
	}
	dst := output + "/" + filename

	if IsExist(dst) {
		dstsize := FileSize(dst)
		if !cover && FileSize(src) == dstsize {
			log.Printf("the same filesize [%d] , ignore [%s] \n", dstsize, dst)
			ch <- filename
			return nil
		}
	}
	size, err := copy(src, dst)
	if err != nil {
		panic(err)
	}
	log.Printf("copy [%s/%s] to [%s/%s] , size:%d \n", input, filename, output, filename, size)
	ch <- filename
	return err
}

func fileNames(input string) []string {
	input = strings.TrimRight(input, "/")
	if !IsReadable(input) {
		panic(fmt.Errorf("[%s] is not readable\n", input))
	}
	//code := fmt.Sprintf("ls %s", input)
	//data, err := cmd.RunCommandOutputString(code)
	data, err := GetDirFileNames(input)
	if err != nil {
		panic(err)
	}
	return data
}

func GetDirFileNames(src string) ([]string, error) {
	rd, err := ioutil.ReadDir(src)
	data := make([]string, 0, 1000)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
			//GetAllFile(src + fi.Name() + "\\")
		}
		data = append(data, fi.Name())
	}
	return data, err
}

func checkPath(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("output is empty\n")
	}
	if !IsExist(path) {
		return fmt.Errorf("[%s] is not exists\n", path)
	}
	return nil
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func IsReadable(name string) bool {
	err := syscall.Access(name, syscall.O_RDONLY)
	if err == nil {
		return true
	}
	return false
}

func IsWriteable(name string) bool {
	err := syscall.Access(name, syscall.O_RDWR)
	if err == nil {
		return true
	}
	return false
}

func FileSize(filename string) int64 {
	if info, err := os.Stat(filename); err == nil {
		return info.Size()
	}

	return 0
}
