// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package main

import "io/ioutil"
import "fmt"
import "strings"
import "os"
import "flag"
import "errors"
import "path/filepath"
import "./githubapi"

func setFlags() (help bool, description, filename string, args []string) {
	flag.BoolVar(&help, "help", false, "show help message")
	flag.StringVar(&filename, "name", "untitled.txt", "filename to send to github")
	flag.StringVar(&description, "desc", "No Description", "description of file")

	flag.Parse()
	args = flag.Args()
	return
}

func printHelpMessage() {
	fmt.Println("gistit v0.2")
	fmt.Println("Reads on stdin for input") // or reads filenames on cmd line
	fmt.Println("Optional Arguments:")
	fmt.Println("\t-help Help Screen")
	fmt.Println("\t-name Name of file on github")
	fmt.Println("\t-desc Description of file")
}

func readStdin() string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading from stdin")
		os.Exit(0)
	}
	return string(bytes)
}

func dummy() {
}

func readFiles(filenames []string) ([]githubapi.File, error) {
	var files []githubapi.File
	for _, f := range filenames {
		c, e := ioutil.ReadFile(f)
		if e != nil {
			fmt.Println("Error reading file", e)
			os.Exit(1)
		}
		cleaned := githubapi.EscapeNLs(string(c))
		files = append(files, githubapi.NewFile(f, cleaned))
	}
	return files, nil
}

func gistitrc() error {
	env := os.Environ()
	homedir, err := getHOMEfromEnv(env)
	if err == nil {
		gistitrcPath := filepath.Join(homedir, ".gistitrc")
		file, err := os.Open(gistitrcPath)
		if err != nil {
			// `~/.gistitrc` does not exit, not that it matters
			return errors.New("No gistitrc file found")
		} else {
			fmt.Println("~/.gistitrc Found, but not supported")
			defer file.Close()
			return nil // someday update return values to get needed auth
		}
		fmt.Println(gistitrcPath)
	}
	fmt.Printf("No $HOME was found")
	return nil
}

func getHOMEfromEnv(env []string) (string, error) {
	homePath := ""
	for i := range env {
		if strings.HasPrefix(env[i], "HOME=") {
			homePath = env[i][5:]
			break
		}
	}
	if homePath == "" {
		return "", errors.New("no gistrc file was found")
	}
	return homePath, nil
}
