
# go-prompt

 Terminal prompts for Go. A fork of the go-prompt package by segmentio.

 https://github.com/segmentio/go-prompt

 View the [docs](http://godoc.org/pkg/github.com/wyattjoh/go-prompt).

## Example

```go
package main

import "github.com/wyattjoh/go-prompt"

var langs = []string{
  "c",
  "c++",
  "lua",
  "go",
  "js",
  "ruby",
  "python",
}

func main() {
  i := prompt.Choose("What's your favorite language?", langs)
  println("picked: " + langs[i])
}
```

## License

 MIT
