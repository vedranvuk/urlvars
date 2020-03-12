// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package urlvars

import "testing"

func TestExpander(t *testing.T) {

	type testitem struct {
		rawurl, template, expected string
	}

	testdata := []testitem{
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"{scheme}{userinfo}{hostname}{port}{path}{query}{fragment}",
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
		},
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"1{scheme}2{userinfo}3{hostname}4{port}5{path}6{query}7{fragment}8",
			"1https://2user:pass@3www.example.com4:805/users/vedran/file.ext6?action=view&mode=quick7#top8",
		},
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"{bogus}{hostname}{bogus}",
			"{bogus}www.example.com{bogus}",
		},
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"{bogus{hostname}{bogus}",
			"{boguswww.example.com{bogus}",
		},
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"{{{bogus}{hostname}{bogus}",
			"{{{bogus}www.example.com{bogus}",
		},
		testitem{
			"https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top",
			"{bogus}{{hostname}}{bogus}",
			"{bogus}{www.example.com}{bogus}",
		},
	}

	for _, data := range testdata {
		if output, err := Expand(data.template, data.rawurl); err != nil {
			t.Fatalf("expand '%s' to '%s' failed: %v\n", data.rawurl, data.template, err)
		} else {
			if output != data.expected {
				t.Fatalf("Expand() failed, want '%s', got '%s'", data.expected, output)
			}
		}
	}

}

func BenchmarkExpander(b *testing.B) {

	const (
		rawurl   = "https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top"
		template = "{scheme}{userinfo}{hostname}{port}{path}{query}{fragment}"
	)

	for i := 0; i < b.N; i++ {
		Expand(template, rawurl)
	}

}
