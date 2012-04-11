package main

import "io/ioutil"
import "fmt"
import "strings"
import "os"
import "flag"
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
	fmt.Println("gistit v0.1")
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

func newlineToNewLine(input string) string {
	a := strings.Replace(input, "\n", "\\n", -1)
	b := strings.Replace(a, "\r\n", "\\n", -1)
	return b
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
		files = append(files, githubapi.NewFile(f, string(c)))
	}
	return files, nil
}
