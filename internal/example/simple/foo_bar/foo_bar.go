package foobar

import (
	"fmt"

	"github.com/gpbPiazza/archp/internal/example/simple/baar"
	"github.com/gpbPiazza/archp/internal/example/simple/foo"
)

func OrchestrateFooAndBaar(fooEx, baarEx string) {
	foo := foo.New(fooEx)
	baar := baar.New(baarEx)

	fmt.Println(foo.Example)
	fmt.Println(baar.Example)
}
