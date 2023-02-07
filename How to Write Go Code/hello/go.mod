module example/user/hello

go 1.19

require (
	github.com/Sun2yKid/Go/Module-Import/myremotepackage v0.0.0-20230207130349-6f8dbff640f1
	github.com/google/go-cmp v0.5.9
	mylocalpackage v0.0.0-00010101000000-000000000000
)

replace mylocalpackage => ../mylocalpackage
