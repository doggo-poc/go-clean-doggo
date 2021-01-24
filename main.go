package main

import (
	"DoggosPkg/doggos/adapters"
	"DoggosPkg/doggos/router"
	"DoggosPkg/doggos/usecase"
	"DoggosPkg/repositories"

	breedsMapper "DoggosPkg/breeds/adapter"
	breedsRouter "DoggosPkg/breeds/router"
	breedsUseCase "DoggosPkg/breeds/usecase"
	breedsRepository "DoggosPkg/repositories/breeds"

	"github.com/labstack/echo"
)

func main() {

	mapper := adapters.NewDoggoMapper()
	usecase := usecase.NewDoggoUseCase(repositories.NewDoggoRepository(), mapper)

	e := echo.New()
	router.NewDoggosHandler(e, usecase)

	breedsUseCase := breedsUseCase.NewBreedsUseCase(breedsRepository.NewBreedsRepository(), breedsMapper.NewBreedsMapper())
	breedsRouter.NewBreedsHandler(e, breedsUseCase)

	e.Logger.Fatal(e.Start(":9090"))
}
