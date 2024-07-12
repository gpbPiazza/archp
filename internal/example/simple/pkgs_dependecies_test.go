package simple

import (
	"testing"

	"github.com/gpbPiazza/archp"
)

func TestDependencies_Foo_Test(t *testing.T) {
	fooAnalizer := archp.NewAnalizer("github.com/gpbPiazza/archp/internal/example/simple/foo")

	err := fooAnalizer.
		DisallowedDependOn("github.com/gpbPiazza/archp/internal/example/simple/baar").
		DisallowedDependOn("github.com/gpbPiazza/archp/internal/example/simple/foo_baar").
		Analize()

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
