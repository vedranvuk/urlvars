# urlvars

Package urlvars implements utilities for parsing elements of an URL into named values.

For now, only path parsing.

WIP.

## Usage

### Path

```
const (
	template = "https://www.example.com/:root/:dir/:subdir/:file"
	rawurl   = "https://www.example.com/home/vedran/temp/file.ext"
)

vars, _ := Path(template, rawurl)
fmt.Printf("%#v\n", vars)

// Outputs:
map[string]string{"dir":"vedran", "file":"file.ext", "root":"home", "subdir":"temp"}

```

## Status

WIP

## License

See included LICENSE.