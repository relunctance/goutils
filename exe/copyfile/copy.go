package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/relunctance/goutils/dump"
	"github.com/relunctance/goutils/fc"
)

const (
	DEFAULT_CHAN_NUM = 20
	MAX_CHAN_NUM     = 100
)

var input string
var output string
var cover bool
var maxnum int
var version string

func main() {
	startTime := time.Now()
	flag.StringVar(&input, "version", "1.0.0.1001", "copy version")
	flag.StringVar(&input, "input", "", "the copy input dir , the type should be dir")
	flag.StringVar(&output, "output", "", "the copy input dir, the type should be dir")
	flag.BoolVar(&cover, "cover", false, "is cover same file ,if false and filesize is equal then will not cover same file (default false)")
	flag.IntVar(&maxnum, "maxnum", DEFAULT_CHAN_NUM, "the num size  of channel , max limit is 100")
	flag.Parse()
	fmt.Println("---------------------------------\n")
	fmt.Println("input:", input)
	fmt.Println("output:", output)
	fmt.Println("cover:", cover)
	fmt.Println("maxnum:", maxnum)
	fmt.Println("\n---------------------------------\n")

	if maxnum > MAX_CHAN_NUM {
		maxnum = MAX_CHAN_NUM
	}
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
	sliceNames := fc.SliceChunk(names, maxnum) //每次最多并发maxnum个
	for _, vs := range sliceNames {
		copyByNames(vs)
		//time.Sleep(1 * time.Second) //等待1秒
	}
	log.Printf("all cost time: [%s]\n", time.Now().Sub(startTime).String())
}

func copyByNames(names []string) {
	l := len(names)
	if l == 0 {
		log.Printf("input dir: [%s] is a empty dir\n", input)
		return
	}

	startTime := time.Now()
	ch := make(chan string, l)
	defer close(ch)
	var sizeTotal int64
	for _, name := range names {
		name = strings.TrimRight(name, "\n")

		go func(filename string) {
			size, err := copyFile(filename, input, output, ch)
			sizeTotal += size
			if err != nil {
				panic(err)
			}
		}(name)

	}

	for i := 0; i < l; i++ {
		_ = <-ch
	}
	costTime := time.Now().Sub(startTime)
	sec := costTime.Seconds()
	speedk := float64(sizeTotal) / sec / 1024
	speedM := speedk / 1024

	dump.Printf("file num: [%d] , cost time: [%s] , total Byte:[%d Byte] , total Format:[%s] speed Kb: [ %.3f Kb/s] , speed Mb [%.3f Mb/s]\n", l, costTime.String(), sizeTotal, fc.ByteFormat(float64(sizeTotal)), speedk, speedM)
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

func copyFile(filename string, input, output string, ch chan string) (int64, error) {
	if !fc.IsWriteable(output) {
		return 0, fmt.Errorf("[%s] is not writeable \n", output)
	}
	src := input + "/" + filename
	if !fc.IsExist(src) {
		return 0, fmt.Errorf("not exists [%s]\n", src)
	}
	dst := output + "/" + filename

	if fc.IsExist(dst) {
		dstsize := fc.FileSize(dst)
		if !cover && fc.FileSize(src) == dstsize { // if cover== false and filesize is ok , then will not cover the same file
			log.Printf("the same filesize [%d] , ignore [%s] \n", dstsize, dst)
			ch <- filename
			return 0, nil
		}
	}
	size, err := copy(src, dst)
	if err != nil {
		panic(err)
	}
	log.Printf("copy [%s/%s] to [%s/%s] , size:%d \n", input, filename, output, filename, size)
	ch <- filename
	return size, err
}

func fileNames(input string) []string {
	input = strings.TrimRight(input, "/")
	if !fc.IsReadable(input) {
		panic(fmt.Errorf("[%s] is not readable\n", input))
	}
	//code := fmt.Sprintf("ls %s", input)
	//data, err := cmd.RunCommandOutputString(code)
	data, err := GetDirFileNames(input)
	dump.Println("names length:", len(data))
	if err != nil {
		panic(err)
	}
	sort.Strings(data)
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
	if !fc.IsExist(path) {
		return fmt.Errorf("[%s] is not exists\n", path)
	}
	return nil
}
