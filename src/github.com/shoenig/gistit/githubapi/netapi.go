// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "strconv"
import "strings"
import "encoding/json"
import "net/http"
import "io/ioutil"
import "fmt"
import "os"

// Returns a gist response (unpacked from JSON reply) that is associated
// with id.
func GetGistResponse(id int) GistResponse {
	resp, httperr := http.Get(api_url + "gists/" + strconv.Itoa(id))
	if httperr != nil {
		fmt.Println("Error connecting to api.github.com", httperr)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return jsonToGistResponse(body)
}

// Returns a list of gist responses (unpacked from JSON reply) that is
// associated with a user login.
func ListUserGists(user string) []GistResponse {
	resp, httperr := http.Get(api_url + "users/" + user + "/gists")
	if httperr != nil {
		fmt.Println("Error connecting to api.github.com", httperr)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var gr []GistResponse
	jsonerr := json.Unmarshal(body, &gr)
	if jsonerr != nil {
		fmt.Printf("Error unmarshaling JSON content", jsonerr)
		os.Exit(1)
	}
	return gr
}

// Push a file (or multiple files) to gist.github.com. This is just a
// convenience function that wraps PushMultiGist, which does the work.
func PushGist(description string, files ...File) GistResponse {
	return PushMultiGist(description, files)
}

// Push a file (or multiple files) to gist.github.com. Currently only
// creating anonymous gists is supported. One possible workaround for now
// is to just fork the anonymous gist into your own account.
func PushMultiGist(description string, files []File) GistResponse {
	// TODO: enable setting public to False (needs OAUTH working)
	g := createNewGist(description, files)
	resp, httperr := http.Post(api_url+"gists", "application/json", strings.NewReader(g))
	if httperr != nil {
		fmt.Println("Error posting gist to api.github.com", httperr)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	sbody := string(body)
	cleaned := nullToEmptyString(sbody)
	cleanbytes := []byte(cleaned)
	gr := jsonToGistResponse(cleanbytes)
	return gr
}

func nullToEmptyString(input string) string {
	return strings.Replace(input, "null", "{}", -1)
}

// description is optional, "" implies no description provided
func createNewGist(description string, files []File) string {
	ret := `{"description": "` + description + `", ` +
		`"public": true, ` +
		`"files": {`
	for _, k := range files {
		ret += k.toJSON() + ", "
	}
	// remove trailing ','
	ret = ret[:len(ret)-2]
	ret += `}` // close files:
	ret += `}` // close entire map

	escaped := EscapeNLs(ret)
	return escaped
}
