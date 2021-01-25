package util

import (
	breedModel "DoggosPkg/breeds/usecase/model"
	breedDtoModel "DoggosPkg/repositories/breeds/model"
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
