package githubapi

import "encoding/json"
import "strconv"
import "net/http"
import "fmt"
import "io/ioutil"
import "os"

const api_url = "https://api.github.com/"

// We do not want users to be able to access these fields directly
// (we do not want them handling interface{}),
// but they must be public for the JSON package to use. So we put them
// in an unexported struct, and put the struct in a wrapper, which then
// provides public accessors.
type internal struct {
	Url         string
	Id          string
	Description string
	Public      bool
	User        map[string]interface{}
	Files       map[string]map[string]interface{}
	Comments    int
	HtmlUrl     string
	GitPullUrl  string
	GitPushUrl  string
	CreatedAt   string
	Forks       []interface{}
	History     []interface{}
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

func (g *Gist) Id() string {
	if !g.initialized {
		noinit()
	}
	return g.data.Id
}

func (g *Gist) Description() string {
	if !g.initialized {
		noinit()
	}
	return g.data.Description
}

func (g *Gist) Public() bool {
	if !g.initialized {
		noinit()
	}
	return g.data.Public
}

func (g *Gist) UserLogin() string {
	if !g.initialized {
		noinit()
	}
	return (g.data.User["login"]).(string)
}

func (g *Gist) UserId() int {
	if !g.initialized {
		noinit()
	}
	return (g.data.User["id"]).(int)
}

func (g *Gist) UserAvatarUrl() string {
	if !g.initialized {
		noinit()
	}
	return (g.data.User["avatar_url"]).(string)
}

func (g *Gist) UserGravatarId() string {
	if !g.initialized {
		noinit()
	}
	return (g.data.User["gravatar_id"]).(string)
}

func (g *Gist) UserUrl() string {
	if !g.initialized {
		noinit()
	}
	return (g.data.User["url"]).(string)
}

func (g *Gist) Comments() int {
	if !g.initialized {
		noinit()
	}
	return g.data.Comments
}

func (g *Gist) HtmlUrl() string {
	if !g.initialized {
		noinit()
	}
	return g.data.HtmlUrl
}

func (g *Gist) GitPullUrl() string {
	if !g.initialized {
		noinit()
	}
	return g.data.GitPullUrl
}

func (g *Gist) GitPushUrl() string {
	if !g.initialized {
		noinit()
	}
	return g.data.GitPushRul
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
