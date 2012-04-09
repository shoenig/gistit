package githubapi

// Provides the content of a History of a Gist
type History struct {
	Url           string
	Version       string
	User          User
	Change_Status changestatus
	Committed_At  string
}

type changestatus struct {
	Deletions int
	Additions int
	Total     int
}
