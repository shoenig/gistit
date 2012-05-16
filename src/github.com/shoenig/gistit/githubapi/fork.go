// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

// A structure compatible with github's fork representation
// inside of a gist JSON structure
type Fork struct {
	User       User
	Url        string
	Created_At string
}
