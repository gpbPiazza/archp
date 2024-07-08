package archp

import (
	"fmt"
	"go/build"
)

func NewAnalizer(importPath string) *Analizer {
	return &Analizer{
		ParentPkgImportPath: importPath,
		disallowedDependOn:  make(map[string]bool),
	}
}

type Analizer struct {
	ParentPkgImportPath string
	disallowedDependOn  map[string]bool
}

func (a *Analizer) Analize() error {
	buildPkg, err := build.Import(a.ParentPkgImportPath, "", build.ImportComment)
	if err != nil {
		return fmt.Errorf("failed to import package: %v", err)
	}

	var errs error
	for _, imp := range buildPkg.Imports {
		if a.disallowedDependOn[imp] {
			err := fmt.Errorf("is not allowed package %s depend on %s", a.ParentPkgImportPath, imp)
			errs = fmt.Errorf("disallowed depend on error: %w", err)
		}
	}

	return errs
}

func (a *Analizer) DisallowedDependOn(importPath string) *Analizer {
	a.disallowedDependOn[importPath] = true
	return a
}
