package githubapi

import "testing"
import "strings"

func Test_CreateNewGist(t *testing.T) {
	f := NewFile("file1.txt", "String file contents")
	d := "the description for this gist"
	json := CreateNewGist(d, f)
	if len(json) == 0 {
		t.Error("CreateNewGist failed to create anything")
	}
	if json[0] != '{' || json[len(json)-1] != '}' {
		t.Error("CreateNewGist severe JSON formatting error")
	}
	if !strings.Contains(json, `"description": "the description for this gist",`) {
		t.Error("CreateNewGist JSON error - description")
	}
	if !strings.Contains(json, `"public": True,`) {
		t.Error("CreateNewGist JSON error - public")
	}
	if !strings.Contains(json, `"files": {"file1.txt": { "content": "String file contents"}, }`) {
		t.Error("CreateNewGist JSON error - files")
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
