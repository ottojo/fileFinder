package main

import (
	"path/filepath"
	"os"
	"flag"
	"log"
	"os/exec"
)

var searchString = "word"

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	if f.Size() > 4000000 {
		return nil
	}

	o, err := exec.Command("grep", searchString, path).Output()
	if len(o) == 0 {
		return nil
	} else {
		log.Printf("Found string %s in file %s in path %s\n", searchString, f.Name(), path)
	}

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	searchString = flag.Arg(1)
	err := filepath.Walk(root, visit)
	log.Printf("filepath.Walk() returned %v\n", err)
}
