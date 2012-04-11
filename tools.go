package main

import "io/ioutil"
import "fmt"
import "strings"
import "os"
import "flag"

func setFlags() (help bool, description, filename string) {
	flag.BoolVar(&help, "help", false, "show help message")
	flag.StringVar(&filename, "name", "noname.txt", "filename to send to github")
	flag.StringVar(&description, "description", "no description", "description of file")

	flag.Parse()
	return
}

func printHelpMessage() {
	fmt.Println("gistit v0.0")
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
	return strings.Replace(input, "\n", "\\n", -1)
}

func dummy() {
}

func ReadFiles(filenames []string) ([]string, error) {
	var contents []string
	for f := range filenames {
		if c, e := ioutil.ReadFile(f); e != nil {
			fmt.Println("Error reading file", e)
			os.Exit(1)
		}
		contents = append(contents, c)
	}
	return contents, nil
}
