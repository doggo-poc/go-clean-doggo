package main

import (
	"DoggosPkg/doggos/adapters"
	"DoggosPkg/doggos/router"
	"DoggosPkg/doggos/usecase"
	"DoggosPkg/repositories"

	"github.com/labstack/echo"
)

func main() {

	mapper := adapters.NewDoggoMapper()
	usecase := usecase.NewDoggoUseCase(repositories.NewDoggoRepository(), mapper)

	e := echo.New()
	router.NewDoggosHandler(e, usecase)

	e.Logger.Fatal(e.Start(":9090"))
}
