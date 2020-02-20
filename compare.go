package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	source := "./source"
	target := "./target"

	compare(source, target)
}

func compare(source string, target string) {
	sourceFiles := getFiles(source)
	targetFiles := getFiles(target)

	for i := 0; i < len(sourceFiles); i++ {
		for _, s := range sourceFiles[i] {
			if !strings.Contains(strings.Join(targetFiles[i], " "), s) {
				fmt.Printf("%s NEW\n", s)
			}
		}
		for _, t := range targetFiles[i] {
			if !strings.Contains(strings.Join(sourceFiles[i], " "), t) {
				fmt.Printf("%s DELETED\n", t)
			}
		}
	}
}

func getDir(root string) []string {
	var dir []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dir = append(dir, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return dir
}

func getFiles(root string) [][]string {
	dir := getDir(root)
	files := make([][]string, len(dir))

	i := 0
	for _, d := range dir {
		file, err := ioutil.ReadDir(d)
		if err != nil {
			panic(err)
		}

		if strings.Contains(d, "./source") {
			d = strings.Replace(d, "./source", "", -1)
		}
		if strings.Contains(d, "source") {
			d = strings.Replace(d, "source/", "", -1)
		}
		if strings.Contains(d, "./target") {
			d = strings.Replace(d, "./target", "", -1)
		}
		if strings.Contains(d, "target") {
			d = strings.Replace(d, "target/", "", -1)
		}

		for _, f := range file {
			if !f.IsDir() {
				files[i] = append(files[i], d+"/"+f.Name())
			}
		}
		i++
	}

	return files
}
