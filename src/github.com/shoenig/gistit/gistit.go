// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

// Package main contains the root of the gistit project. It
// makes use of the githubapi package to provide file pushing
// capabilities to github's gist service.
//
// References:
//  http://api.github.com/
package main

import "fmt"
import "os"
import "github.com/shoenig/gistit/githubapi"

func main() {
	dummy() // emacs.el does not format correctly without this
	help, desc, fname, args := setFlags()
	if help != false {
		printHelpMessage()
		os.Exit(0)
	}

	if len(args) == 0 {
		all := readStdin()
		gr := githubapi.PushGist(desc, githubapi.NewFile(fname, all))
		fmt.Println(gr.Html_Url)
	} else {
		files, ferr := readFiles(args)
		if ferr != nil {
			fmt.Println("Error reading file", ferr)
			os.Exit(1)
		}
		gr := githubapi.PushMultiGist(desc, files)
		fmt.Println(gr.Html_Url)
	}
}
