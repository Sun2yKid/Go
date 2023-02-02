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

# How to Write Go Code (with GOPATH)
https://go.dev/doc/gopath_code

SettingGOPATH
https://github.com/golang/go/wiki/SettingGOPATH
> go env -w GOPATH=/Users/zhonghui.yao/go:/Users/zhonghui.yao/MyProjects/GoLang

GOPATH and Modules

When using modules, GOPATH is no longer used for resolving imports.
However, it is still used to store downloaded source code (in GOPATH/pkg/mod)
and compiled commands (in GOPATH/bin).
