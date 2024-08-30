Package archp
=================
[![Go Reference](https://pkg.go.dev/badge/github.com/gpbPiazza/archp.svg)](https://pkg.go.dev/github.com/gpbPiazza/archp)
![License](https://img.shields.io/dub/l/vibe-d.svg)


achp enforce architectural rules as errors values or unit tests

It has the following features:

-   Validate pkgs imports as dependencies.
-   Handles custom errs for each architecture policy

Installation
------------

Use go get.

	go get github.com/gpbPiazza/archp

Then import the archp package into your own code.

	import "github.com/gpbPiazza/archp"


Usage and documentation
------

Please see https://pkg.go.dev/github.com/gpbPiazza/archp/ for detailed usage docs.

##### Examples:

- [Simple](./internal/example/simple/pkgs_dependecies_test.go)


Disclaimers
------
### Special Notes:
1 .archp depends totally on `go/build` package to work, so if your package under test is not in a remote repository it will always return err in the `build.Import` function.
2. This first version doesn't support validate a package children packages imports, for example:
````
-- domain
        |-- services
        |        |-- service.go
        |-- models
        |        |-- model.go
````
The analysis under `domain` package is not supported to validate if the `service.go` file or `model.go` are following the expected rules.

#### Future features:
- solve the especial notes 2. Packages analysis to nested packages dependencies.
- options to use `t *testing.T`

License
-------
Distributed under MIT License, please see license file within the code for more details.

Maintainers
-----------
- [gpbPiazza](https://github.com/gpbPiazza)
