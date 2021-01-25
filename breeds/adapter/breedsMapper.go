package adapter

import (
	"DoggosPkg/breeds/usecase/model"
	dtoModel "DoggosPkg/repositories/breeds/model"
)

type BreedsMapper interface {
	Map(breeds []dtoModel.BreedDto) []model.Breed
}

type breedsMapper struct {
}

func NewBreedsMapper() *breedsMapper {
	return &breedsMapper{}
}

func (mapper *breedsMapper) Map(breeds []dtoModel.BreedDto) []model.Breed {
	var retVal = make([]model.Breed, 0)
	for _, breedDto := range breeds {
		breed := model.Breed{
			ID:          breedDto.Id,
			Name:        breedDto.Name,
			BreedGroup:  breedDto.BreedGroup,
			LifeSpan:    breedDto.LifeSpan,
			Temperament: breedDto.Temperament,
		}
		retVal = append(retVal, breed)
	}
	return retVal
}
