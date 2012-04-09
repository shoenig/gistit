package githubapi

import "testing"

func Test_GetFile_Filename(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.GetFile("ring.erl").Filename != "ring.erl" {
		t.Error("f.Filename incorrect")
	}
}

func Test_GetFile_Size(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.GetFile("ring.erl").Size != 932 {
		t.Error("f.Size incorrect")
	}
}

func Test_GetFile_Raw_Url(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.GetFile("ring.erl").Raw_Url != "https://gist.github.com/raw/365370/8c4d2d43d178df44f4c03a7f2ac0ff512853564e/ring.erl" {
		t.Error("f.Raw_Url incorrect")
	}
}

func Test_GetFile_Content(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	if g.GetFile("ring.erl").Content != "contents of gist" {
		t.Error("f.Content incorrect")
	}
}

func Test_GetFiles(t *testing.T) {
	g := jsonToGist([]byte(GET_GIST))
	files := g.GetFiles()
	if len(files) != 1 {
		t.Error("g.GetFiles length failed")
	}
	if files[0].Filename != "ring.erl" {
		t.Error("g.GetFiles failed")
	}
}
