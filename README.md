# urlvars

Package urlvars implements utilities for parsing elements of an URL into named values.

## Example

### UrlVars

```
 template := https://www.example.com/:root/:sub/:file
 rawurl := https://www.example.com/users/vedran/.listfiles.sh?action=list#listing
 values, err := Path(template, rawurl)
 
 // values will be:
 // map[string]string{"root": "users", "sub": "vedran", "file": ".listfiles.sh"}
```

### Expander

```
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

  rawurl =   https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top
  template = {scheme}{userinfo}{hostname}{port}{path}{query}{fragment}

  Expand(template, exampleurl)

//  Output: https://user:pass@www.example.com:80/users/vedran/file.ext?action=view&mode=quick#top
```

## Documentation

[Godoc](https://godoc.org/github.com/vedranvuk/urlvars)


## License

See included LICENSE.