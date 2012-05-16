// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "testing"

func Test_jsonToGistResponse_url(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Url != "https://api.github.com/gists/1" {
		t.Error("g.Url is incorrect")
	}
}

func Test_jsonToGistResponse_id(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Id != "1" {
		t.Error("g.Id is incorrect")
	}
}

func Test_jsonToGistResponse_description(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Description != "description of gist" {
		t.Error("g.Description is incorrect")
	}
}

func Test_jsonToGistResponse_public(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if !g.Public {
		t.Error("g.Public is incorrect")
	}
}

func Test_jsonToGistResponse_comments(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Comments != 0 {
		t.Error("g.Comments is incorrect")
	}
}

func Test_jsonToGistResponse_html_url(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Html_Url != "https://gist.github.com/1" {
		t.Error("g.Html_Url is incorrect")
	}
}

func Test_jsonToGistResponse_git_push_url(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Git_Pull_Url != "git://gist.github.com/1.git" {
		t.Error("g.Git_Pull_Url is incorrect")
	}
}

func Test_jsonToGistResponse_created_at(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	if g.Created_At != "2010-04-14T02:15:15Z" {
		t.Error("g.Created_At is incorrect")
	}
}
