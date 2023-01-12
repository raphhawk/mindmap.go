package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

var loc = os.Args[1]

func rgrab(f fs.FileInfo, h map[fs.FileInfo]string, baseDir string) {
	if !f.IsDir() {
		h[f] = baseDir + "/" + f.Name()
		return
	}
	dirString, err := ioutil.ReadDir(baseDir + "/" + f.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range dirString {
		rgrab(i, h, baseDir+"/"+f.Name())
	}
}

func main() {
	if len(os.Args) > 2 {
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(loc)
	if err != nil {
		log.Panic(err)
	}
	files2 := map[fs.FileInfo]string{}
	for _, i := range files {
		rgrab(i, files2, loc)
	}
	for _, i := range files2 {
		fmt.Println(i)
	}
}
