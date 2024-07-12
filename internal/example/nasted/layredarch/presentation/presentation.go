package presentation

import (
	"fmt"

	"github.com/gpbPiazza/archp/internal/example/nasted/layredarch/application"
	"github.com/gpbPiazza/archp/internal/example/nasted/layredarch/domain"
	"github.com/gpbPiazza/archp/internal/example/nasted/layredarch/infraestructure/db/nosql/repositores"
)

func Presentations() {
	domain.NewModelService(repositores.NoSQL{})
	fmt.Println(application.MyApp{})
}
