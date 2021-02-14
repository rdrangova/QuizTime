package options

import (
	"fmt"
	"os"
	"path/filepath"
)

var files []string

func list(dirrectory string) []string {

	root := dirrectory
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func printQuizList() {
	for _, file := range files {
		fmt.Println(file[:len(file)-3])
	}
}
