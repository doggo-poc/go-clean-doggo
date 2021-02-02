package util

import (
	petsModel "DoggosPkg/pets/usecase/model"
	doggosDtoModel "DoggosPkg/repositories"
	catsDtoModel "DoggosPkg/repositories/cats/model"
	"fmt"
)

func GenerateMockedPetsBreeds(n int) []petsModel.PetBreed {
	retVal := make([]petsModel.PetBreed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, petsModel.PetBreed{
			Id:          i,
			Name:        fmt.Sprintf("Name %d", i),
			BreedGroup:  fmt.Sprintf("BreedGroup %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GeneratePets(n int) []petsModel.Pet {
	retVal := make([]petsModel.Pet, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, petsModel.Pet{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  fmt.Sprintf("Width %d", i),
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedPetsBreeds(n),
		})
	}
	return retVal
}

func GeneratePetsFromDoggsCats(dDto []doggosDtoModel.DoggoDto, cDto []catsDtoModel.CatDto) []petsModel.Pet {
	retVal := make([]petsModel.Pet, 0)
	if dDto != nil {
		for _, d := range dDto {
			bPet := make([]petsModel.PetBreed, 0)
			for _, b := range d.Breeds {
				bPet = append(bPet, petsModel.PetBreed{
					Id:          b.Id,
					Name:        b.Name,
					BreedGroup:  b.BreedGroup,
					LifeSpan:    b.LifeSpan,
					Temperament: b.Temperament,
				})
			}
			retVal = append(retVal, petsModel.Pet{
				Height: d.Height,
				Id:     d.Id,
				Width:  d.Width,
				Url:    d.Url,
				Breeds: bPet,
			})
		}
	}

	if cDto != nil {
		for _, c := range cDto {
			bPet := make([]petsModel.PetBreed, 0)
			for _, b := range c.Breeds {
				bPet = append(bPet, petsModel.PetBreed{
					Id:          b.Id,
					Name:        b.Name,
					BreedGroup:  b.BreedGroup,
					LifeSpan:    b.LifeSpan,
					Temperament: b.Temperament,
				})
			}
			retVal = append(retVal, petsModel.Pet{
				Height: c.Height,
				Id:     c.Id,
				Width:  c.Width,
				Url:    c.Url,
				Breeds: bPet,
			})
		}
	}
	return retVal
}
