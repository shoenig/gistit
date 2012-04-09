package githubapi

// File provides methods which provide actual content and some
// meta information about the individual files that make up a gist.
type File struct {
	Name    string
	Size    int
	Raw_url string
	Content string
}
