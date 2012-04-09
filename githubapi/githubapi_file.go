package githubapi

// Provides content of individual files that make up a gist
type File struct {
	Filename string
	Size     int
	Raw_Url  string
	Content  string
}
