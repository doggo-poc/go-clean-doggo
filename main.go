package main

import (
	"DoggosPkg/doggos/adapters"
	"DoggosPkg/doggos/router"
	"DoggosPkg/doggos/usecase"
	"DoggosPkg/middleware"
	"DoggosPkg/repositories"

	breedsMapper "DoggosPkg/breeds/adapter"
	breedsRouter "DoggosPkg/breeds/router"
	breedsUseCase "DoggosPkg/breeds/usecase"
	breedsRepository "DoggosPkg/repositories/breeds"

	catsMapper "DoggosPkg/cats/adapter"
	catsRouter "DoggosPkg/cats/router"
	catsUseCase "DoggosPkg/cats/usecase"
	catsRepository "DoggosPkg/repositories/cats"

	petsMapper "DoggosPkg/pets/adapter"
	petsRouter "DoggosPkg/pets/router"
	petsUseCase "DoggosPkg/pets/usecase"

	"github.com/labstack/echo"
)

func main() {

	mapper := adapters.NewDoggoMapper()
	usecase := usecase.NewDoggoUseCase(repositories.NewDoggoRepository(), mapper)

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	router.NewDoggosHandler(e, usecase)

	breedsUseCase := breedsUseCase.NewBreedsUseCase(breedsRepository.NewBreedsRepository(), breedsMapper.NewBreedsMapper())
	breedsRouter.NewBreedsHandler(e, breedsUseCase)

	catRepo := catsRepository.NewCatRepository()
	catsUseCase := catsUseCase.NewCatUseCase(catRepo, catsMapper.NewCatMapper())
	catsRouter.NewCatsHandler(e, catsUseCase)

	pm := petsMapper.NewPetMapper(adapters.NewDoggoMapper(), catsMapper.NewCatMapper())
	dr := repositories.NewDoggoRepository()
	cr := catsRepository.NewCatRepository()
	pr := petsUseCase.NewPetsUseCase(dr, cr, pm)
	petsRouter.NewPetsHandler(e, pr)
	e.Logger.Fatal(e.Start(":9090"))
}
