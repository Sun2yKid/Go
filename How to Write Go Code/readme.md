# How to Write Go Code
https://go.dev/doc/code

## Code organization

Go programs are organized into packages. 

A *package* is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A *module* is a collection of related Go packages that are released together. A file named go.mod there declares the *module path*: the import path prefix for all packages within the module.

An *import path* is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. 

> go list -m all

The command go list -m all lists the current module and all its dependencies

> go list -m -versions module

list the available tagged versions of that module

> go install example/user/hello  | go install . | go install

The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).


You can use the go env command to portably set the default value for an environment variable for future go commands:
> go env -w GOBIN=/somewhere/else/bin

To unset a variable previously set by go env -w, use go env -u:
> go env -u GOBIN


For added convenience, we'll add the install directory to our PATH to make running binaries easy:
```
# Windows users should consult https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```

The `go mod tidy` command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.


Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:

> go clean -modcache
