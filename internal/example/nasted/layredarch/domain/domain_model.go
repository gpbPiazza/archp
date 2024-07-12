package domain

import (
	"fmt"

	anypkg1 "github.com/gpbPiazza/archp/internal/example/nasted/layredarch/infraestructure/any_pkg_1"
)

type querier interface {
	Query()
}

type MyModelService struct {
	querier querier
}

func NewModelService(querier querier) MyModelService {
	return MyModelService{
		querier: querier,
	}
}

func (mm MyModelService) DoSomething() {
	mm.querier.Query()
	fmt.Println(anypkg1.PKG1{})
}
