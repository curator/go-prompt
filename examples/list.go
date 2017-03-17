package main

import prompt "github.com/wyattjoh/go-prompt"

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
