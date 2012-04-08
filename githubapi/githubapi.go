package githubapi

import "encoding/json"
import "strconv"
import "net/http"

//import "errors"
import "fmt"
import "io/ioutil"
import "os"

const api_url = "https://api.github.com/"


/* We do not want users to be able to access these fields directly,
but they must be public for the JSON package to use. So we put them
in an unexported struct, and put the struct in a wrapper, which then
provides public accessors */
type internal struct {
	Url          string
	Id           string
	Description  string
	Public       bool
	User         map[string]interface{}
	Files        map[string]map[string]interface{}
	Comments     int
	Html_url     string
	Git_pull_url string
	Git_push_url string
	Created_at   string
	Forks        []interface{}
	History      []interface{}
}

func noinit() {
	panic("cannot call method on un-initialized gist")
}

type Gist struct {
	data        internal
	initialized bool
}

func (g Gist) String() string {
	if !g.initialized {
		noinit()
	}
	return "Gist!"
}

func (g *Gist) Url() string {
	if !g.initialized {
		noinit()
	}
	return g.data.Url
}

func (g *Gist) setInternal(i internal) {
	if g.initialized {
		noinit()
	}
	g.initialized = true
	g.data = i
}

func GetGist(id int) Gist {
	resp, httperr := http.Get(api_url + "gists/" + strconv.Itoa(id))
	if httperr != nil {
		fmt.Println("Error connecting to api.github.com", httperr)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var i internal
	jsonerr := json.Unmarshal(body, &i)
	if jsonerr != nil {
		fmt.Printf("Error unmarshaling JSON content", jsonerr)
		os.Exit(1)
	}
	var g Gist
	g.setInternal(i)
	return g
}
