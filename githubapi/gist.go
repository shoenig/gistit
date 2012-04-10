package githubapi

import "encoding/json"
import "strconv"
import "net/http"
import "fmt"
import "io/ioutil"
import "os"

type Gist struct {
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

func (g Gist) String() string {
	return "TODO"
}

func (g *Gist) UserLogin() string {
	return g.UserMap["user"].Login
}

func (g *Gist) UserId() int {
	return g.UserMap["user"].Id
}

func (g *Gist) UserAvatarUrl() string {
	return g.UserMap["user"].Avatar_Url
}

func (g *Gist) UserGravatarId() string {
	return g.UserMap["user"].Gravatar_Id
}

func (g *Gist) UserUrl() string {
	return g.UserMap["user"].Url
}

func (g *Gist) GetFiles() []File {
	var files []File
	for _, v := range g.Files {
		files = append(files, v)
	}
	return files
}

func (g *Gist) GetFile(filename string) File {
	return g.Files[filename]
}

func jsonToGist(jsondata []byte) Gist {
	var g Gist
	jsonerr := json.Unmarshal(jsondata, &g)
	if jsonerr != nil {
		fmt.Printf("Error unmarshaling JSON content", jsonerr)
		os.Exit(1)
	}
	return g
}

func GetGist(id int) Gist {
	resp, httperr := http.Get(api_url + "gists/" + strconv.Itoa(id))
	if httperr != nil {
		fmt.Println("Error connecting to api.github.com", httperr)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return jsonToGist(body)
}
