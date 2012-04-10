package githubapi

import "strconv"
import "encoding/json"
import "net/http"
import "io/ioutil"
import "fmt"
import "os"

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

// description is optional, "" implies no description provided
func CreateNewGist(description string, files ...File) string {
	ret := "{\"description\": \"" + description + "\"," +
		"\"public\": True," +
		"\"files\": {"
	for _, k := range files {
		ret += k.toJSON() + ", "
	}
	ret += "}" // close files:
	ret += "}" // close entire map
	return ret
}
