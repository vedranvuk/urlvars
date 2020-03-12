// Copyright 2019 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package urlvars

import (
	"net/url"
	"strings"
)

// expand expands a known key enclosed in single curly braces or returns the key as is.
func expand(key string, url *url.URL) (value string, expanded bool) {
	expanded = true
	switch strings.ToLower(key) {
	case "{scheme}":
		value = url.Scheme + "://"
	case "{userinfo}":
		value = url.User.String() + "@"
	case "{host}":
		value = url.Host
	case "{hostname}":
		value = url.Hostname()
	case "{port}":
		value = ":" + url.Port()
	case "{path}":
		value = url.Path
	case "{query}":
		value = "?" + url.RawQuery
	case "{fragment}":
		value = "#" + url.Fragment
	default:
		value = key
		expanded = false
	}
	return
}

// Expand expands the rawurl according to template.
//
// URL keys are enclosed in single curly braces. Non-paired braces,
// extra braces and unrecognized keys in braces are passed to output as-is.
//
// Given an example url:
//
//  https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top
//
// The following supported keys would return the following values:
//
//  scheme:   scheme part of URL, "https://"
//  userinfo: userinfo part of URL, "user:pass@"
//  host:     host part of URL, "www.example.com:80"
//  hostname: host part of URL, "www.example.com"
//  port:     port part of URL, ":80"
//  path:     path part of URL, "/users/vedran/file.ext"
//  query:    query part of URL, "?action=view&mode=quick"
//  fragment: query part of URL, "#top"
//
// Example:
//
//  rawurl =   https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top
//  template = {scheme}{userinfo}{hostname}{port}{path}{query}{fragment}
//
//  Expand(template, exampleurl)
//  Output: https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top
func Expand(template, rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", ErrParse.WrapCause("invalid rawurl", err)
	}
	output := ""
restart:
	for len(template) > 0 {
		for left := 0; left < len(template); left++ {
			if template[left] == '{' {
				for right := left; right < len(template); right++ {
					if template[right] == '}' {
						if value, expanded := expand(template[left:right+1], url); expanded {
							output += template[:left] + value
							if right+1 == len(template) {
								template = ""
							} else {
								template = template[right+1:]
							}
							goto restart
						} else {
							break
						}
					}
				}
			}
		}
		output += template
		template = ""
	}
	return output, nil
}
