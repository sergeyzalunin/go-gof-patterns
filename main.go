package main

import "github.com/sergeyzalunin/go-gof-patterns/builder"


func main() {
	b := builder.NewHTMLBuilder("main")
	print(b.String())
}