Tutorial: Getting started with multi-module workspaces

https://go.dev/doc/tutorial/workspaces

In this tutorial, you’ll create two modules in a shared multi-module workspace, make changes across those modules, and see the results of those changes in a build.



> go work init ./hello
The go work init command tells go to create a go.work file for a workspace containing the modules in the ./hello directory.

> go work use ./example

go work use [-r] [dir] adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.

go work edit edits the go.work file similarly to go mod edit
go work sync syncs dependencies from the workspace’s build list into each of the workspace modules.



"go.work" can be used instead of adding replace directives to work across multiple modules.