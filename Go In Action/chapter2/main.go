package main

import (
	"log"
	"os"

	_ "chapter2/matchers" // 下划线是为了让Go对包做init操作，但不使用包里的标识符(类型，函数，常量，接口)。
	"chapter2/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
