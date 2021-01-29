package adapter

import (
	"DoggosPkg/doggos/usecase/model"
	petsModel "DoggosPkg/pets/usecase/model"
	"DoggosPkg/repositories"
	catsModelDto "DoggosPkg/repositories/cats/model"
)

type PetsMapper interface {
	Map(d []repositories.DoggoDto, c []catsModelDto.CatDto) []petsModel.Pet
}

type petsMapper struct {
}

func NewPetMapper() *petsMapper {
	return &petsMapper{}
}

func mapBreedsDto(breedsDto []repositories.BreedDto) []model.Breed {
	var retVal = make([]model.Breed, 0)
	for _, breedDto := range breedsDto {
		breed := model.Breed{
			Id:          breedDto.Id,
			Name:        breedDto.Name,
			BreedGroup:  breedDto.BreedGroup,
			LifeSpan:    breedDto.LifeSpan,
			Temperament: breedDto.Temperament,
		}
		retVal = append(retVal, breed)
	}
	return retVal
}

func mapDoggoDto(doggoDto repositories.DoggoDto) model.Doggo {
	return model.Doggo{
		Height: doggoDto.Height,
		Width:  doggoDto.Width,
		Id:     doggoDto.Id,
		Url:    doggoDto.Url,
		Breeds: mapBreedsDto(doggoDto.Breeds),
	}
}

func (mapper *petsMapper) Map(d []repositories.DoggoDto, c []catsModelDto.CatDto) []petsModel.Pet {
	var retVal = make([]petsModel.Pet, 0)
	if d != nil {
		for _, doggoDto := range d {
			retVal = append(retVal, doggoToPet(doggoDto))
		}
	}
	if c != nil {
		for _, catDto := range c {
			retVal = append(retVal, catToPet(catDto))
		}
	}
	return retVal
}

func doggoToPet(doggo repositories.DoggoDto) petsModel.Pet {
	var breeds = make([]petsModel.PetBreed, 0)
	for _, breedDto := range doggo.Breeds {
		breed := petsModel.PetBreed{
			Id:          breedDto.Id,
			Name:        breedDto.Name,
			BreedGroup:  breedDto.BreedGroup,
			LifeSpan:    breedDto.LifeSpan,
			Temperament: breedDto.Temperament,
		}
		breeds = append(breeds, breed)
	}

	return petsModel.Pet{
		Height: doggo.Height,
		Width:  doggo.Width,
		Id:     doggo.Id,
		Url:    doggo.Url,
		Breeds: breeds,
	}
}

func catToPet(cat catsModelDto.CatDto) petsModel.Pet {
	var breeds = make([]petsModel.PetBreed, 0)
	for _, breedDto := range cat.Breeds {
		breed := petsModel.PetBreed{
			Id:          breedDto.Id,
			Name:        breedDto.Name,
			BreedGroup:  breedDto.BreedGroup,
			LifeSpan:    breedDto.LifeSpan,
			Temperament: breedDto.Temperament,
		}
		breeds = append(breeds, breed)
	}

	return petsModel.Pet{
		Height: cat.Height,
		Width:  cat.Width,
		Id:     cat.Id,
		Url:    cat.Url,
		Breeds: breeds,
	}
}
