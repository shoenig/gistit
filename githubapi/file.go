package githubapi

// Provides content of individual files that make up a gist
type File struct {
	Filename string
	Size     int
	Raw_Url  string
	Content  string
}

func NewFile(fname, content string) File {
	return File{fname, 0, "", content}
}

// Return github's expected JSON format for a file that is to be
// uploaded as a part of a gist. The end-user does not use this,
// instead they should be using CreateNewGist, which creates the
// entire JSON structure for uploading to github
func (f *File) toJSON() string {
	// SAMPLE FORMAT:
	//  "file1.txt": {
	//     "content": "String file contents"
	//  }
	ret := "\"" + f.Filename + "\": { \"content\": \"" + f.Content +
		"\"}"
	return ret
}
