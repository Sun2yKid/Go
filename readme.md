# Tutorials

https://go.dev/doc/tutorial/

# The Go Playground
https://go.dev/play/


# How to Write Go Code
https://go.dev/doc/code

A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository.

An import path is a string used to import a package. 


# How to Write Go Code (with GOPATH)
https://go.dev/doc/gopath_code

SettingGOPATH
https://github.com/golang/go/wiki/SettingGOPATH
> go env -w GOPATH=/Users/zhonghui.yao/go:/Users/zhonghui.yao/MyProjects/GoLang

GOPATH and Modules

When using modules, GOPATH is no longer used for resolving imports.
However, it is still used to store downloaded source code (in GOPATH/pkg/mod)
and compiled commands (in GOPATH/bin).
