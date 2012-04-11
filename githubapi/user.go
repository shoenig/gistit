// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

// Provides content of user type
type User struct {
	Login       string
	Id          int
	Avatar_Url  string
	Gravatar_Id string
	Url         string
}
