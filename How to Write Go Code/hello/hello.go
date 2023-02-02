package main

import (
	"fmt"

	"example/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println("How to Write Go Code: Hello, world.")
	fmt.Println(morestrings.ReverseRunes("How to Write Go Code: Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
