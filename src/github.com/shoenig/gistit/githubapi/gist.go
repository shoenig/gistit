// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "encoding/json"
import "fmt"
import "os"

// A GistResponse is designed to be a compatible representation
// with github's JSON representation of a Gist. Go's builtin encoding/json
// package makes use of this struct when unpacking data recieved from
// github.
type GistResponse struct {
	Url          string
	Id           string
	Description  string
	Public       bool
	UserMap      map[string]User
	Files        map[string]File
	Comments     int
	Html_Url     string
	Git_Pull_Url string
	Git_Push_Url string
	Created_At   string
	Forks        []Fork
	History      []History
}

// A String representation of a JSON response from github.
func (g GistResponse) String() string {
	bytes, marshalerr := json.Marshal(g)
	if marshalerr != nil {
		return ""
	}
	return string(bytes)
}

// The user login provided in a JSON response from github.
func (g *GistResponse) UserLogin() string {
	return g.UserMap["user"].Login
}

// The user id provided in a JSON response from github.
func (g *GistResponse) UserId() int {
	return g.UserMap["user"].Id
}

// The user avatar url provided in a JSON response from github.
func (g *GistResponse) UserAvatarUrl() string {
	return g.UserMap["user"].Avatar_Url
}

// The user gravatar id provided in a JSON response from github.
func (g *GistResponse) UserGravatarId() string {
	return g.UserMap["user"].Gravatar_Id
}

// The user url provided in a JSON response from github.
func (g *GistResponse) UserUrl() string {
	return g.UserMap["user"].Url
}

// The list of files provided in a JSON response from github.
func (g *GistResponse) GetFiles() []File {
	var files []File
	for _, v := range g.Files {
		files = append(files, v)
	}
	return files
}

// The specific file associated with filename that is provided in
// a JSON response from github.
func (g *GistResponse) GetFile(filename string) File {
	return g.Files[filename]
}

func jsonToGistResponse(jsondata []byte) GistResponse {
	var g GistResponse
	jsonerr := json.Unmarshal(jsondata, &g)
	if jsonerr != nil {
		fmt.Printf("Error unmarshaling JSON content", jsonerr)
		fmt.Println(string(jsondata))
		os.Exit(1)
	}
	return g
}
