package githubapi

import "strconv"
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
