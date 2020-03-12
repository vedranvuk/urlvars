// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package urlvars

import (
	"errors"
	"testing"
)

func TestURLVars(t *testing.T) {

	type testitem struct {
		rawurl, template string

		values map[string]string
		err    error
	}

	testdata := []testitem{
		testitem{
			"https://www.example.com/home/vedran/temp/file.ext",
			"https://www.example.com/:root/:dir/:subdir/:file",
			map[string]string{"root": "home", "dir": "vedran", "subdir": "temp", "file": "file.ext"},
			nil,
		},
		testitem{
			"https://www.example.com/home/vedran/temp/file.ext",
			"https://www.example.com/:root/:root/:root/:root",
			map[string]string{},
			ErrDupKey,
		},
		testitem{
			"foo:/boo?baz*bar#gaz//bogus",
			"https://www.example.com/:root/:root/:root/:root",
			map[string]string{},
			ErrParse,
		},
		testitem{
			"https://www.example.com/home/vedran/temp/file.ext",
			"bogus",
			map[string]string{},
			ErrParse,
		},
		testitem{
			"https://www.example.com/one/two/",
			"https://www.example.com/:one/:two/:three/",
			map[string]string{},
			ErrParse,
		},
		testitem{
			"https://www.example.com/one/two/three/",
			"https://www.example.com/:one/:two/",
			map[string]string{},
			ErrParse,
		},
		testitem{
			"https://www.example.com/one/two/three/",
			"https://www.example.com/::one/:two:/::three::/",
			map[string]string{":one": "one", "two:": "two", ":three::": "three"},
			nil,
		},
	}

	for _, testitem := range testdata {
		vars, err := Path(testitem.template, testitem.rawurl)
		if !errors.Is(err, testitem.err) {
			t.Fatalf("expand '%s' to '%s' failed: want error '%v', got '%v'\n", testitem.rawurl, testitem.template, testitem.err, err)
		}
		for expkey, expval := range testitem.values {
			if resval, ok := vars[expkey]; ok {
				if resval != expval {
					t.Fatalf("expand '%s' to '%s' failed: variable '%s' has wrong value, want '%s', got '%s'\n",
						testitem.rawurl, testitem.template, expkey, expval, resval)
				}
			} else {
				t.Fatalf("expand '%s' to '%s' failed: variable '%s' not parsed\n",
					testitem.rawurl, testitem.template, expkey)
			}
		}
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
