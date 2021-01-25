package usecase

import (
	"DoggosPkg/breeds/adapter"
	"DoggosPkg/breeds/usecase/model"
	"DoggosPkg/repositories/breeds"
)

type BreedsUseCase interface {
	GetBreeds() ([]model.Breed, error)
}

type breedsUseCase struct {
	repository breeds.BreedsRepository
	mapper     adapter.BreedsMapper
}

func NewBreedsUseCase(breedsRepo breeds.BreedsRepository, mapper adapter.BreedsMapper) BreedsUseCase {
	return &breedsUseCase{
		repository: breedsRepo,
		mapper:     mapper,
	}
}

func (b *breedsUseCase) GetBreeds() ([]model.Breed, error) {
	breeds, err := b.repository.GetBreeds()
	if err != nil {
		return nil, err
	}
	return b.mapper.Map(breeds), nil
}
