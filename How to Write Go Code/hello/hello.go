package main

import (
	"fmt"

	"example/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/Sun2yKid/Go/Module-Import/myremotepackage/mypackage"
)

func main() {
    fmt.Println("How to Write Go Code: Hello, world.")
	fmt.Println(morestrings.ReverseRunes("How to Write Go Code: Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	mypackage.New()
}
