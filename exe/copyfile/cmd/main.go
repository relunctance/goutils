package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/relunctance/goutils/exe/copyfile"
)

var input string
var output string
var cover bool
var maxnum int
var version string

func main() {
	flag.StringVar(&input, "version", "1.0.0.1001", "copy version")
	flag.StringVar(&input, "input", "", "the copy input dir , the type should be dir")
	flag.StringVar(&output, "output", "", "the copy input dir, the type should be dir")
	flag.BoolVar(&cover, "cover", false, "is cover same file ,if false and filesize is equal then will not cover same file (default false)")
	flag.IntVar(&maxnum, "maxnum", copyfile.DEFAULT_CHAN_NUM, "the num size  of channel , max limit is 100")
	flag.Parse()
	fmt.Println("---------------------------------\n")
	fmt.Println("input:", input)
	fmt.Println("output:", output)
	fmt.Println("cover:", cover)
	fmt.Println("maxnum:", maxnum)
	fmt.Println("\n---------------------------------\n")
	input = strings.TrimRight(input, "/")
	output = strings.TrimRight(output, "/")
	if input == "" || output == "" {
		flag.Usage()
		return
	}
	copyfile.Copy(input, output, maxnum, cover)

}
