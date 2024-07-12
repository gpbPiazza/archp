package archp

import (
	"errors"
	"testing"
)

func TestPKGAnalizer_Constructors(t *testing.T) {
	t.Run("return pkg analizer with disallowedDependOn prop not nil", func(t *testing.T) {
		pkgAnalizer := NewAnalizer("testing")

		if pkgAnalizer.disallowedDependOn == nil {
			t.Error("expected to have disallowedDependOn not nil")
		}
	})

	t.Run("return pkg analizer with ImportPath prop with the same passaed as argument", func(t *testing.T) {
		pkgAnalized := "testing"
		pkgAnalizer := NewAnalizer(pkgAnalized)

		if pkgAnalizer.importPath != pkgAnalized {
			t.Errorf("expected to have ImportPath equals to argument passed but got: %s", pkgAnalizer.importPath)
		}
	})

	t.Run("return pkg analizer with disallowedDependOn map with fmt key", func(t *testing.T) {
		dependOnPkg := "fmt"
		pkgAnalizer := NewAnalizer("testing").DisallowedDependOn(dependOnPkg)

		if !pkgAnalizer.disallowedDependOn[dependOnPkg] {
			t.Errorf("expected to have disallowedDependOn key %s", dependOnPkg)
		}
	})

	t.Run("return disallowedDependOn map with many keys", func(t *testing.T) {
		dependOnPkg1 := "fmt"
		dependOnPkg2 := "net/http"

		pkgAnalizer := NewAnalizer("testing").
			DisallowedDependOn(dependOnPkg1).
			DisallowedDependOn(dependOnPkg2)

		if !pkgAnalizer.disallowedDependOn[dependOnPkg1] {
			t.Errorf("expected to have disallowedDependOn key %s", dependOnPkg1)
		}

		if !pkgAnalizer.disallowedDependOn[dependOnPkg2] {
			t.Errorf("expected to have disallowedDependOn key %s", dependOnPkg2)
		}
	})

	t.Run("if DisallowedDependOn is called twice with the same value should not change the key value", func(t *testing.T) {
		dependOnPkg1 := "fmt"

		pkgAnalizer := NewAnalizer("testing").
			DisallowedDependOn(dependOnPkg1)

		if !pkgAnalizer.disallowedDependOn[dependOnPkg1] {
			t.Errorf("expected to have disallowedDependOn key %s", dependOnPkg1)
		}

		pkgAnalizer = pkgAnalizer.DisallowedDependOn(dependOnPkg1)

		if !pkgAnalizer.disallowedDependOn[dependOnPkg1] {
			t.Errorf("expected to have disallowedDependOn key %s", dependOnPkg1)
		}
	})
}

func TestPKGAnalizer_Analize(t *testing.T) {
	t.Run("return error on import package when has a poor importpath", func(t *testing.T) {
		analizer := NewAnalizer("poor/import/path")

		err := analizer.Analize()

		if !errors.Is(err, ErrImport) {
			t.Errorf("expected to receive %v, but got %v", ErrImport, err)
		}
	})

	t.Run("return error of dissallowed to depend on when pkg analized fails on analize", func(t *testing.T) {
		analizer := NewAnalizer("github.com/gpbPiazza/archp").DisallowedDependOn("errors")

		err := analizer.Analize()

		expectedErr := &PolicyError{
			TargetAnalized: "github.com/gpbPiazza/archp",
			Policy:         policyDissalowedDependOn,
			TriggerErr:     "errors",
		}

		if !errors.As(err, &expectedErr) {
			t.Errorf("expected to receive %v error, but got %v", expectedErr, err)
		}
	})

	t.Run("return many errors of dissallowed to depend on when pkg analized fails on analize", func(t *testing.T) {
		dissalowedDependOnPKGs := []string{"errors", "fmt", "go/build"}

		analizer := NewAnalizer("github.com/gpbPiazza/archp")

		var expectedErrs []error
		for _, pkg := range dissalowedDependOnPKGs {
			analizer = analizer.DisallowedDependOn(pkg)

			errFromPKG := newPolicyError("github.com/gpbPiazza/archp", policyDissalowedDependOn, pkg)
			expectedErrs = append(expectedErrs, errFromPKG)
		}
		expectedErr := errors.Join(expectedErrs...)

		err := analizer.Analize()

		if err.Error() != expectedErr.Error() {
			t.Errorf("expected to have %v but got %v", expectedErr, err)
		}
	})

	t.Run("return no errors when all policies are ok", func(t *testing.T) {
		analizer := NewAnalizer("github.com/gpbPiazza/archp").
			DisallowedDependOn("github.com/gpbPiazza/archp/internal").
			DisallowedDependOn("github.com/gpbPiazza/archp/internal/example")

		err := analizer.Analize()

		if err != nil {
			t.Errorf("expected to err to be nil but got %v", err)
		}
	})

}
