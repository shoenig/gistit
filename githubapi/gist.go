package githubapi

import "encoding/json"
import "strconv"
import "net/http"
import "fmt"
import "io/ioutil"
import "os"

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

func (g GistResponse) String() string {
	bytes, marshalerr := json.Marshal(g)
	if marshalerr != nil {
		return ""
	}
	return string(bytes)
}

func (g *GistResponse) UserLogin() string {
	return g.UserMap["user"].Login
}

func (g *GistResponse) UserId() int {
	return g.UserMap["user"].Id
}

func (g *GistResponse) UserAvatarUrl() string {
	return g.UserMap["user"].Avatar_Url
}

func (g *GistResponse) UserGravatarId() string {
	return g.UserMap["user"].Gravatar_Id
}

func (g *GistResponse) UserUrl() string {
	return g.UserMap["user"].Url
}

func (g *GistResponse) GetFiles() []File {
	var files []File
	for _, v := range g.Files {
		files = append(files, v)
	}
	return files
}

func (g *GistResponse) GetFile(filename string) File {
	return g.Files[filename]
}

func jsonToGistResponse(jsondata []byte) GistResponse {
	var g GistResponse
	jsonerr := json.Unmarshal(jsondata, &g)
	if jsonerr != nil {
		fmt.Printf("Error unmarshaling JSON content", jsonerr)
		os.Exit(1)
	}
	return g
}

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
