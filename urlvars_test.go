// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package urlvars

import (
	"testing"
)

func TestURLVars(t *testing.T) {

	const (
		template = "https://www.example.com/:root/:dir/:subdir/:file"
		rawurl   = "https://www.example.com/home/vedran/temp/file.ext"
	)

	vars, err := Path(template, rawurl)
	if err != nil {
		t.Fatal(err)
	}

	if vars["root"] != "home" {
		t.Fatal("fail")
	}
	if vars["dir"] != "vedran" {
		t.Fatal("fail")
	}
	if vars["subdir"] != "temp" {
		t.Fatal("fail")
	}
	if vars["file"] != "file.ext" {
		t.Fatal("fail")
	}
}

func BenchmarkURLVars(b *testing.B) {

	const (
		template = "https://www.example.com/:root/:dir/:subdir/:file"
		rawurl   = "https://www.example.com/home/vedran/temp/file.ext?action=test"
	)

	for i := 0; i < b.N; i++ {
		Path(template, rawurl)
	}
}
