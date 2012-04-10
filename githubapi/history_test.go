package githubapi

import "testing"

func Test_GetHistories(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	hists := g.History
	if len(hists) != 1 {
		t.Error("g.History is wrong length")
	}
}

func Test_History_Url(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	h := g.History[0]
	if h.Url != "https://api.github.com/gists/1/57a7f021a713b1c5a6a199b54cc514735d2d462f" {
		t.Error("History Url is incorrect")
	}
}

func Test_History_Version(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	h := g.History[0]
	if h.Version != "57a7f021a713b1c5a6a199b54cc514735d2d462f" {
		t.Error("History Version is incorrect")
	}
}

func Test_History_User(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	h := g.History[0]
	if h.User.Login != "octocat" {
		t.Error("History User.Login incorrect")
	}
	if h.User.Id != 1 {
		t.Error("History User.Id incorrect")
	}
	if h.User.Avatar_Url != "https://github.com/images/error/octocat_happy.gif" {
		t.Error("History User.Avatar_Url incorrect")
	}
	if h.User.Gravatar_Id != "somehexcode" {
		t.Error("History User.Gravatar_Id incorrect")
	}
	cs := h.Change_Status
	if cs.Deletions != 0 {
		t.Error("History.Change_Status is incorrect")
	}
	if cs.Additions != 180 {
		t.Error("History.Change_Status is incorrect")
	}
	if cs.Total != 180 {
		t.Error("History.Change_Status is incorrect")
	}
}
