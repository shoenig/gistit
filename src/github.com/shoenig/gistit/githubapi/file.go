// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "path/filepath"

// Provides content of individual files that make up a gist
type File struct {
	Filename string
	Size     int
	Raw_Url  string
	Content  string
}

// Create a new File, which is compatible with github's
// JSON representation of a file within a gist.
func NewFile(fname, content string) File {
	return File{fname, 0, "", content}
}

// Return github's expected JSON format for a file that is to be
// uploaded as a part of a gist. The end-user does not use this,
// instead they should be using CreateNewGist, which creates the
// entire JSON structure for uploading to github. The filename
// will be transformed to the base of the actual (local) filename.
// Ex: `/tmp/foo.txt` will be named `foo.txt`
func (f *File) toJSON() string {
	// SAMPLE FORMAT:
	//   "file1.txt": {"content": "String file contents"}
	n := filepath.Base(f.Filename)
	c := EscapeNLs(f.Content)
	ret := "\"" + n + "\": { \"content\": \"" +
		c + "\"}"
	return ret
}
