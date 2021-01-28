package adapter

import (
	"DoggosPkg/cats/usecase/model"
	catsDto "DoggosPkg/repositories/cats/model"
)

type CatsMapper interface {
	Map(doggos []catsDto.CatDto) []model.Cat
}

type catsMapper struct {
}

func NewCatMapper() *catsMapper {
	return &catsMapper{}
}

func mapBreedsDto(breedsDto []catsDto.CatBreedDto) []model.CatBreed {
	var retVal = make([]model.CatBreed, 0)
	for _, breedDto := range breedsDto {
		breed := model.CatBreed{
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

func mapCatDto(catDto catsDto.CatDto) model.Cat {
	return model.Cat{
		Height: catDto.Height,
		Width:  catDto.Width,
		Id:     catDto.Id,
		Url:    catDto.Url,
		Breeds: mapBreedsDto(catDto.Breeds),
	}
}

func (mapper *catsMapper) Map(cats []catsDto.CatDto) []model.Cat {
	var retVal = make([]model.Cat, 0)
	for _, catDto := range cats {
		retVal = append(retVal, mapCatDto(catDto))
	}
	return retVal
}
