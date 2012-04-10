package githubapi

import "testing"
import "strings"

func Test_createNewGist(t *testing.T) {
	f := NewFile("file1.txt", "String file contents")
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
	if !strings.Contains(json, `"public": True,`) {
		t.Error("createNewGist JSON error - public")
	}
	if !strings.Contains(json, `"files": {"file1.txt": { "content": "String file contents"}, }`) {
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
