package githubapi

import "testing"

func Test_jsonToGist_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Url != "https://api.github.com/gists/1" {
		t.Error("g.Url is incorrect")
	}
}

func Test_jsonToGist_id(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Id != "1" {
		t.Error("g.Id is incorrect")
	}
}

func Test_jsonToGist_description(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Description != "description of gist" {
		t.Error("g.Description is incorrect")
	}
}

func Test_jsonToGist_public(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if !g.Public {
		t.Error("g.Public is incorrect")
	}
}

func Test_jsonToGist_comments(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Comments != 0 {
		t.Error("g.Comments is incorrect")
	}
}

func Test_jsonToGist_html_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Html_Url != "https://gist.github.com/1" {
		t.Error("g.Html_Url is incorrect")
	}
}

func Test_jsonToGist_git_push_url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Git_Pull_Url != "git://gist.github.com/1.git" {
		t.Error("g.Git_Pull_Url is incorrect")
	}
}

func Test_jsonToGit_created_at(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.Created_At != "2010-04-14T02:15:15Z" {
		t.Error("g.Created_At is incorrect")
	}
}
