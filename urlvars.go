// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package urlvars implements utilities for parsing elements of an URL into named values.
package urlvars

import (
	"net/url"
	"strings"

	"github.com/vedranvuk/errorex"
)

var (
	// ErrURLVars is the base error of urlvars package.
	ErrURLVars = errorex.New("urlvars")
	// ErrParse is a parse error.
	ErrParse = ErrURLVars.Wrap("parse error")
	// ErrDupKey is returned when a duplicate key is specified.
	ErrDupKey = ErrURLVars.WrapFormat("duplicate key '%s'")
	// ErrInvalidTemplate is returned when an invalid template is specified.
	ErrInvalidTemplate = ErrURLVars.Wrap("invalid template")
)

// parsepath extracts path from a raw url and splits it on elements.
func parsepath(rawurl string) ([]string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return strings.Split(u.Path, "/")[1:], nil
}

// Path extracts specific path elements into a map of named variables.
//
// Path elements are marked as variables by prefixing an element with a colon.
//
// Example:
//
//  template := https://www.example.com/:root/:sub/:file
//  rawurl := https://www.example.com/users/vedran/.listfiles.sh?action=list#listing
//
// returns a map with following values:
//  {"root": "users", "sub": "vedran", "file": ".listfiles.sh"}
//
// Extra path elements not matched by template are ignored.
//
// If an error occurs it is returned with a nil map.
func Path(template, rawurl string) (map[string]string, error) {

	tmplelems, err := parsepath(template)
	if err != nil {
		return nil, ErrParse.WrapCause("invalid template url", err)
	}

	rawelems, err := parsepath(rawurl)
	if err != nil {
		return nil, ErrParse.WrapCause("invalid raw url", err)
	}

	if len(tmplelems) == 0 {
		return nil, ErrInvalidTemplate
	}

	m := make(map[string]string)
	rawcount := len(rawelems)
	for idx, val := range tmplelems {
		if idx >= rawcount {
			break
		}
		if strings.HasPrefix(val, ":") && val != rawelems[idx] && len(val) > 1 {
			if _, exists := m[val[1:]]; exists {
				return nil, ErrDupKey.WrapArgs(m[val[1:]])
			}
			m[val[1:]] = rawelems[idx]
		}
	}

	return m, nil
}
