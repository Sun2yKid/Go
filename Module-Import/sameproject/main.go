package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	// "time"
	"sameproject/mypackage"

	"github.com/google/go-cmp/cmp"
)

type Duration int64


func main() {
	mypackage.New()
    fmt.Println("Hello, World!")
	//time.Duration
	var dur Duration
	dur = int64(100)

	r, err := http.Get()
	r.Body

	os.Stdout

	io.Copy()

	slice := make([]int, 2, 3)


	fmt.Println(cmp.Diff("Hello Go!", "Hello Go"))
}
