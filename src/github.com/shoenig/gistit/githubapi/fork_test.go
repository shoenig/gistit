// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "testing"

func Test_GetForks(t *testing.T) {
	g := jsonToGistResponse([]byte(GET_GIST))
	forks := g.Forks
	if len(forks) != 1 {
		t.Error("GetForks is wrong length")
	}
	if forks[0].User.Login != "octocat" {
		t.Error("GetForks[0] is not filled in correctly")
	}
	if forks[0].Url != "https://api.github.com/gists/5" {
		t.Error("GetForks[0] Url is incorrect")
	}
	if forks[0].Created_At != "2011-04-14T16:00:49Z" {
		t.Error("GetForks[0] Created_At is incorrect")
	}
}
