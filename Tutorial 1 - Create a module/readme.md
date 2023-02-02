https://go.dev/doc/tutorial/create-module

Go code is grouped into packages, and packages are grouped into modules.

To enable dependency tracking for your code by creating a go.mod file, run
> go mod init module_name

- The go build command compiles the packages, along with their dependencies, but it doesn't install the results.
- The go install command compiles and installs the packages.


You can discover the install path by running the go list command, as in the following example:
> go list -f '{{.Target}}'

Add the go install directory to your system's shell path
- On linux or mac
> export PATH=$PATH:/path/to/your/install/directory

- On windows
> set PATH=%PATH%;C:\path\to\your\install\directory

As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

> $ go env -w GOBIN=/path/to/your/bin

or in Windows OS

> $ go env -w GOBIN=C:\path\to\your\bin

unset a variable
> go env -u GOBIN

check go env
> $ go env

install the program
> go install "module name"

This command builds the hello command, producing an executable binary. It then installs that binary as $HOME/go/bin/hello (or, under Windows, %USERPROFILE%\go\bin\hello.exe).

The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).

check go version
> $ go version
