package foobar

import (
	"fmt"

	"github.com/gpbPiazza/archp/internal/example/baar"
	"github.com/gpbPiazza/archp/internal/example/foo"
)

func OrchestrateFooAndBaar(fooEx, baarEx string) {
	foo := foo.New(fooEx)
	baar := baar.New(baarEx)

	fmt.Println(foo.Example)
	fmt.Println(baar.Example)
}
