package example

import (
	"testing"

	"github.com/gpbPiazza/archp"
)

func TestDependencies_Foo_Test(t *testing.T) {
	fooAnalizer := archp.NewPKGAnalizer("github.com/gpbPiazza/archp/internal/example/foo")

	err := fooAnalizer.
		DisallowedDependOn("github.com/gpbPiazza/archp/internal/example/baar").
		DisallowedDependOn("github.com/gpbPiazza/archp/internal/example/foo_baar").
		Analize()

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
