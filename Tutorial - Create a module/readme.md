https://go.dev/doc/tutorial/create-module

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

or

> $ go env -w GOBIN=C:\path\to\your\bin

check go env
> $ go env

check go version
> $ go version
