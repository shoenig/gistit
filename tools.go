package main

import "io/ioutil"
import "errors"
import "fmt"
import "os"

func ReadFiles(filenames []string) ([]string, error) {
	var contents []string
	for f := range filenames {
		if c, e := ioutil.ReadFile(f); e != nil {
			fmt.Println("Error reading file", e)
			os.Exit(1)
		}
		contents = append(contents, c)
	}
	return contents
}
