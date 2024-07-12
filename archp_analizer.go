package archp

import (
	"errors"
	"fmt"
	"go/build"
	"strings"
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
		disallowedDependOn: make([]string, 0),
	}
}

type Analizer struct {
	importPath         string
	disallowedDependOn []string
}

func (a *Analizer) Analize() error {
	pkg, err := build.Import(a.importPath, localImport, build.ImportComment)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrImport, err)
	}

	var errs []error
	for _, importPath := range pkg.Imports {
		errs = append(errs, a.analizeDisallowedDependOn(importPath))
	}

	return errors.Join(errs...)
}

func (a *Analizer) DisallowedDependOn(importPath ...string) *Analizer {
	a.disallowedDependOn = append(a.disallowedDependOn, importPath...)
	return a
}

func (a *Analizer) analizeDisallowedDependOn(importPath string) error {
	for _, dependOn := range a.disallowedDependOn {
		if strings.HasPrefix(importPath, dependOn) {
			return newPolicyError(a.importPath, policyDissalowedDependOn, importPath)
		}
	}
	return nil
}
