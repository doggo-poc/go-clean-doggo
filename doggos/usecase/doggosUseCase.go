package usecase

import (
	"DoggosPkg/doggos/adapters"
	"DoggosPkg/doggos/usecase/model"
	"DoggosPkg/repositories"
)

type DoggoUseCase interface {
	GetDoggos(page int, limit int, breedID string) ([]model.Doggo, error)
}

type doggoUseCase struct {
	repository repositories.DoggoRepository
	mapper     adapters.DoggosMapper
}

func NewDoggoUseCase(doggoRepo repositories.DoggoRepository, doggosMapper adapters.DoggosMapper) DoggoUseCase {
	return &doggoUseCase{
		repository: doggoRepo,
		mapper:     doggosMapper,
	}
}

func (d *doggoUseCase) GetDoggos(page int, limit int, breedID string) ([]model.Doggo, error) {
	doggos, err := d.repository.GetDoggos(page, limit, breedID)
	if err != nil {
		return nil, err
	}
	return d.mapper.Map(doggos), nil
}
