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

const GET_GIST = `{
  "url": "https://api.github.com/gists/1",
  "id": "1",
  "description": "description of gist",
  "public": true,
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "somehexcode",
    "url": "https://api.github.com/users/octocat"
  },
  "files": {
    "ring.erl": {
      "size": 932,
      "filename": "ring.erl",
      "raw_url": "https://gist.github.com/raw/365370/8c4d2d43d178df44f4c03a7f2ac0ff512853564e/ring.erl",
      "content": "contents of gist"
    }
  },
  "comments": 0,
  "html_url": "https://gist.github.com/1",
  "git_pull_url": "git://gist.github.com/1.git",
  "git_push_url": "git@gist.github.com:1.git",
  "created_at": "2010-04-14T02:15:15Z",
  "forks": [
    {
      "user": {
        "login": "octocat",
        "id": 1,
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "somehexcode",
        "url": "https://api.github.com/users/octocat"
      },
      "url": "https://api.github.com/gists/5",
      "created_at": "2011-04-14T16:00:49Z"
    }
  ],
  "history": [
    {
      "url": "https://api.github.com/gists/1/57a7f021a713b1c5a6a199b54cc514735d2d462f",
      "version": "57a7f021a713b1c5a6a199b54cc514735d2d462f",
      "user": {
        "login": "octocat",
        "id": 1,
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "somehexcode",
        "url": "https://api.github.com/users/octocat"
      },
      "change_status": {
        "deletions": 0,
        "additions": 180,
        "total": 180
      },
      "committed_at": "2010-04-14T02:15:15Z"
    }
  ]
}`
