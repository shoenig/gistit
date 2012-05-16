// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "strings"

const api_url = "https://api.github.com/"

func EscapeNLs(input string) string {
	a := strings.Replace(input, "\r\n", "\\n", -1)
	b := strings.Replace(a, "\n", "\\n", -1)
	return b
}
