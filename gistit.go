package main

import "fmt"
import "os"
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
