package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

var loc = os.Args[1]

type Node struct {
	data     []fs.FileInfo
	nxtNodes []Node
}

func (dirs *Node) stretch() {
	var s = []fs.FileInfo{dirs.data[0]}
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		if cur.IsDir() {
			var n Node
			var err error
			n.data, err = ioutil.ReadDir(loc + "/" + cur.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			dirs.nxtNodes = append(dirs.nxtNodes, n)
		}
	}
}

func grab(loc string) (Node, error) {
	files, err := ioutil.ReadDir(loc)
	if err != nil {
		return Node{}, err
	}
	dirs := Node{
		files,
		[]Node{},
	}
	dirs.stretch()
	return dirs, nil
}

func main() {
	files, err := grab(loc)
	if err != nil {
		log.Panic(err)
	}
	for _, i := range files.data {
		fmt.Println(i.Name(), i.IsDir())
	}
}
