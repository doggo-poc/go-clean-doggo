package util

import (
	breedModel "DoggosPkg/breeds/usecase/model"
	catsModel "DoggosPkg/cats/usecase/model"
	breedDtoModel "DoggosPkg/repositories/breeds/model"
	catsDtoModel "DoggosPkg/repositories/cats/model"
	"fmt"
)

func GenerateMockedBreedDto(n int) []breedDtoModel.BreedDto {
	retVal := make([]breedDtoModel.BreedDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, breedDtoModel.BreedDto{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedBreed(n int) []breedModel.Breed {
	retVal := make([]breedModel.Breed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, breedModel.Breed{
			ID:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedCatBreedDto(n int) []catsDtoModel.CatBreedDto {
	retVal := make([]catsDtoModel.CatBreedDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsDtoModel.CatBreedDto{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateCatBreed(n int) []catsModel.CatBreed {
	retVal := make([]catsModel.CatBreed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsModel.CatBreed{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedCatsDto(n int) []catsDtoModel.CatDto {
	retVal := make([]catsDtoModel.CatDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsDtoModel.CatDto{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  fmt.Sprintf("Width %d", i),
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedCatBreedDto(n),
		})
	}
	return retVal
}

func GenerateCat(n int) []catsModel.Cat {
	retVal := make([]catsModel.Cat, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsModel.Cat{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  fmt.Sprintf("Width %d", i),
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateCatBreed(n),
		})
	}
	return retVal
}
