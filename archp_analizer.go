package archp

import (
	"errors"
	"fmt"
	"go/build"
)

var (
	localImport = "."

	// ErrImport is the base err to any err returned from the go/build pkg.
	// Use ErrImport to assert against go/Build errs.
	ErrImport = errors.New("failed to import package")

	policyDissalowedDependOn = "dissalowed depend on policy"
)

func NewPKGAnalizer(importPath string) *PKGAnalizer {
	return &PKGAnalizer{
		importPath:         importPath,
		disallowedDependOn: make(map[string]bool),
	}
}

type PKGAnalizer struct {
	importPath         string
	disallowedDependOn map[string]bool
}

func (pkg *PKGAnalizer) Analize() error {
	buildPkg, err := build.Import(pkg.importPath, localImport, build.ImportComment)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrImport, err)
	}

	var errs []error
	for _, importPath := range buildPkg.Imports {
		if pkg.disallowedDependOn[importPath] {
			errs = append(errs, newPolicyError(pkg.importPath, policyDissalowedDependOn, importPath))
		}
	}

	return errors.Join(errs...)
}

func (pkg *PKGAnalizer) DisallowedDependOn(importPath string) *PKGAnalizer {
	pkg.disallowedDependOn[importPath] = true
	return pkg
}
