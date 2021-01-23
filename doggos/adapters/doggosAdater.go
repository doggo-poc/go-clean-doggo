package adapters

import (
	"DoggosPkg/doggos/usecase/model"
	"DoggosPkg/repositories"
)

type DoggosMapper interface {
	Map(doggos []repositories.DoggoDto) []model.Doggo
}

type doggosMapper struct {
}

func NewDoggoMapper() *doggosMapper {
	return &doggosMapper{}
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
		Id:     doggoDto.Id,
		Url:    doggoDto.Url,
		Breeds: mapBreedsDto(doggoDto.Breeds),
	}
}

func (mapper *doggosMapper) Map(doggos []repositories.DoggoDto) []model.Doggo {
	var retVal = make([]model.Doggo, 0)
	for _, doggoDto := range doggos {
		retVal = append(retVal, mapDoggoDto(doggoDto))
	}
	return retVal
}
