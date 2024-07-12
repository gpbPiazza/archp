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

func NewAnalizer(importPath string) *Analizer {
	return &Analizer{
		importPath:         importPath,
		disallowedDependOn: make(map[string]bool),
	}
}

type Analizer struct {
	importPath         string
	disallowedDependOn map[string]bool
}

func (a *Analizer) Analize() error {
	pkg, err := build.Import(a.importPath, localImport, build.ImportComment)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrImport, err)
	}

	var errs []error
	for _, importPath := range pkg.Imports {
		if a.disallowedDependOn[importPath] {
			errs = append(errs, newPolicyError(a.importPath, policyDissalowedDependOn, importPath))
		}
	}

	return errors.Join(errs...)
}

func (pkg *Analizer) DisallowedDependOn(importPath string) *Analizer {
	pkg.disallowedDependOn[importPath] = true
	return pkg
}
