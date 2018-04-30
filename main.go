package main

import (
	"path/filepath"
	"os"
	"flag"
	"log"
	"os/exec"
)

var rootFolder string
var searchString string
var maxFileSize int64
var minFileSize int64

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	if f.Size() > maxFileSize || f.Size() < minFileSize {
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
	flag.Int64Var(&maxFileSize, "maxFileSize", 5000000, "Maximum File Size in bytes")
	flag.Int64Var(&minFileSize, "minFileSize", 1, "Minimum File Size in bytes")
	flag.StringVar(&searchString, "searchString", "word", "String to search for")
	flag.StringVar(&rootFolder, "rootFolder", ".", "Root search folder")
	flag.Parse()

	err := filepath.Walk(rootFolder, visit)
	if err != nil {
		log.Panic(err)
	}
}
