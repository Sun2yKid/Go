# Tutorials
https://go.dev/doc/tutorial/

# The Go Playground
https://go.dev/play/


# How to Write Go Code
https://go.dev/doc/code


# Refer
See Effective Go for tips on writing clear, idiomatic Go code.

https://go.dev/doc/effective_go

Visit the documentation page for a set of in-depth articles about the Go language and its libraries and tools.

https://go.dev/doc/#articles

## SettingGOPATH
https://github.com/golang/go/wiki/SettingGOPATH
> go env -w GOPATH=/Users/username/go

## Set GOBIN
You can use the go env command to portably set the default value for an environment variable for future go commands:
> go env -w GOBIN=/Users/username/go/bin

To unset a variable previously set by go env -w, use go env -u:
> go env -u GOBIN

Add to OS env PATH
> export PATH=$PATH:/Users/username/go/bin

## goimports
install
> go install golang.org/x/tools/cmd/goimports@latest

usage
> goimports -w filename.go

# How to Write Go Code (with GOPATH)
https://go.dev/doc/gopath_code


GOPATH and Modules

When using modules, GOPATH is no longer used for resolving imports.
However, it is still used to store downloaded source code (in GOPATH/pkg/mod)
and compiled commands (in GOPATH/bin).
