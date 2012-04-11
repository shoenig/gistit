package main

import "errors"
import "fmt"
import "os"
import "path/filepath"
import "strings"
import "./githubapi"

func main() {
	dummy() // emacs.el does not format correctly without this
	help, desc, fname, args := setFlags()
	if help != false {
		printHelpMessage()
		os.Exit(0)
	}

	if len(args) == 0 {
		all := readStdin()
		cleaned := newlineToNewLine(all)
		gr := githubapi.PushGist(desc, githubapi.NewFile(fname, cleaned))
		fmt.Println(gr.Html_Url)
	} else {
		files, ferr := readFiles(args)
		if ferr != nil {
			fmt.Println("Error reading file", ferr)
			os.Exit(1)
		}
		gr := githubapi.PushMultiGist(desc, files)
		fmt.Println(gr)
		fmt.Println(gr.Html_Url)
	}
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
