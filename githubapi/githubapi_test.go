package githubapi

import "testing"

func Test_jsonToGist_initialized(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if !g.initialized {
		t.Error("gist was not initialized upon being filled in")
	}
}

func Test_jsonToGist_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.data.Url != "https://api.github.com/gists/1" {
		t.Error("g.data.Url is incorrect")
	}
	if g.Url() != "https://api.github.com/gists/1" {
		t.Error("g.Url() is incorrect")
	}
}

func Test_jsonToGist_id(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.data.Id != "1" {
		t.Error("g.data.Id is incorrect")
	}
	if g.Id() != "1" {
		t.Error("g.Id() is incorrect")
	}
}

func Test_jsonToGist_description(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.data.Description != "description of gist" {
		t.Error("g.data.Description is incorrect")
	}
	if g.Description() != "description of gist" {
		t.Error("g.Description() is incorrect")
	}
}

func Test_jsonToGist_public(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if !g.data.Public {
		t.Error("g.data.Public is incorrect")
	}
	if !g.Public() {
		t.Error("g.Public() is incorrect")
	}
}

func Test_jsonToGist_user(t *testing.T) {
}

func Test_jsonToGist_files(t *testing.T) {
}

func Test_jsonToGist_comments(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.data.Comments != 0 {
		t.Error("g.data.Comments is incorrect")
	}
	if g.Comments() != 0 {
		t.Error("g.Comments() is incorrect")
	}
}

func Test_jsonToGist_html_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))

	if g.data.HtmlUrl != "https://gist.github.com/1" {
		t.Error("g.data.HtmlUrl is incorrect, %v", g.data.HtmlUrl)
	}
	if g.HtmlUrl() != "https://gist.github.com/1" {
		t.Error("g.HtmlUrl() is incorrect, %v", g.HtmlUrl())
	}
}

func Test_jsonToGist_git_push_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.data.GitPullUrl != "git://gist.github.com/1.git" {
		t.Error("g.data.GitPullUrl is incorrect, %v", g.data.GitPullUrl)
	}
	if g.HtmlUrl() != "git://gist.github.com/1.git" {
		t.Error("g.HtmlUrl() is incorrect, %v", g.HtmlUrl())
	}
}