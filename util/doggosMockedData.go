package util

import (
	doggosModel "DoggosPkg/doggos/usecase/model"
	doggosDto "DoggosPkg/repositories"
	"fmt"
)

func GenerateMockedDogsBreedDto(n int) []doggosDto.BreedDto {
	retVal := make([]doggosDto.BreedDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, doggosDto.BreedDto{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedDogsDto(n int) []doggosDto.DoggoDto {
	retVal := make([]doggosDto.DoggoDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, doggosDto.DoggoDto{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  fmt.Sprintf("Width %d", i),
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedDogsBreedDto(n),
		})
	}
	return retVal
}

func GenerateMockedDogsBreeds(n int) []doggosModel.Breed {
	retVal := make([]doggosModel.Breed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, doggosModel.Breed{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateDoggos(n int) []doggosModel.Doggo {
	retVal := make([]doggosModel.Doggo, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, doggosModel.Doggo{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  fmt.Sprintf("Width %d", i),
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedDogsBreeds(n),
		})
	}
	return retVal
}
