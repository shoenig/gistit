// Copyright 2012 Seth Hoenig. All rights reserved.
// Use of this source code is goverened by a GPL-style
// license that can be found in the LICENSE file.

package githubapi

import "testing"

func Test_EscapeNLs_empty_string(t *testing.T) {
	if EscapeNLs("") != "" {
		t.Error("ENLs failed on empty string")
	}
}

func Test_EscapeNLs_single(t *testing.T) {
	if EscapeNLs("\n") != "\\n" {
		t.Error("ENLs failed on single unix nl")
	}
	if EscapeNLs("\r\n") != "\\n" {
		t.Error("ENLs failed on single windows nl")
	}
}

func Test_EscapeNLs_basic(t *testing.T) {
	if EscapeNLs("foo\nbar") != "foo\\nbar" {
		t.Error("ENLs failed on mixed unix nl")
	}
	if EscapeNLs("foo\r\nbar") != "foo\\nbar" {
		t.Error("ENLs failed on mix windows nl")
	}
	a := EscapeNLs("foo\n\r\nbar")
	b := "foo\\n\\nbar"
	if a != b {
		t.Error("ENLs failed on unix/windows nl mix")
	}
}
