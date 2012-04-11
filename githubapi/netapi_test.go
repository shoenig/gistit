// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "testing"
import "strings"

func Test_createNewGist(t *testing.T) {
	f := []File{NewFile("file1.txt", "String file contents")}
	d := "the description for this gist"
	json := createNewGist(d, f)
	if len(json) == 0 {
		t.Error("createNewGist failed to create anything")
	}
	if json[0] != '{' || json[len(json)-1] != '}' {
		t.Error("createNewGist severe JSON formatting error")
	}
	if !strings.Contains(json, `"description": "the description for this gist",`) {
		t.Error("createNewGist JSON error - description")
	}
	if !strings.Contains(json, `"public": true,`) {
		t.Error("createNewGist JSON error - public")
	}
	if !strings.Contains(json, `"files": {"file1.txt": { "content": "String file contents"}}`) {
		t.Error("createNewGist JSON error - files")
	}
}

func Test_createNewGistThreeFiles(t *testing.T) {
	f1 := NewFile("file1.txt", "Contents of File 1!")
	f2 := NewFile("file2.txt", "Contents of File 2.")
	f3 := NewFile("file3.txt", "Contents of File 3?")
	fs := []File{f1, f2, f3}
	d := "A Gist made of THREE files!!"
	json := createNewGist(d, fs)
	if len(json) == 0 {
		t.Error("createNewGist failed to create anything")
	}
	if json[0] != '{' || json[len(json)-1] != '}' {
		t.Error("createNewGist severe JSON formatting error")
	}
	if !strings.Contains(json, `"description": "A Gist made of THREE files!!"`) {
		t.Error("createNewGist JSON error - description")
	}

	if !strings.Contains(json, `"public": true,`) {
		t.Error("createNewGist JSON error - public")
	}
	if !strings.Contains(json, `"files": {"file1.txt": { "content": "Contents of File 1!"}, "file2.txt": { "content": "Contents of File 2."}, "file3.txt": { "content": "Contents of File 3?"}}`) {
		t.Error("createNewGist JSON error - files")
	}
}

/*  SAMPLE PATTERN:
{
  "description": "the description for this gist",
  "public": true,
  "files": {
    "file1.txt": {
      "content": "String file contents"
    }
  }
}
*/
