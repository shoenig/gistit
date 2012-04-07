package main

import "errors"
import "fmt"
import "io/ioutil"
import "net/http"
import "os"
import "path/filepath"
import "strings"

func main() {
	gistitrcOK := gistitrc()
	if gistitrcOK == nil {
		fmt.Println("Using .gistitrc")
	} else {
		fmt.Println("Not using .gistitrc")
	}

	getGist(1)
}

func getGist(id int) {
	resp, err := http.Get("https://api.github.com/gists/1")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Ok", resp)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
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