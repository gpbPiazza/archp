package layredarch

import (
	"testing"

	"github.com/gpbPiazza/archp"
)

func TestOpenLayredArchitecture(t *testing.T) {
	// The Domain layer or packages children of domain
	// should not depend on packages or chidlren packages from the packages
	// application and presentation
	notAllowedDependOnLayers := []string{
		"github.com/gpbPiazza/archp/internal/example/nasted/layredarch/application",
		"github.com/gpbPiazza/archp/internal/example/nasted/layredarch/presentation",
	}

	analizer := archp.
		NewAnalizer("github.com/gpbPiazza/archp/internal/example/nasted/layredarch/domain").
		DisallowedDependOn(notAllowedDependOnLayers...)

	err := analizer.Analize()

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}
