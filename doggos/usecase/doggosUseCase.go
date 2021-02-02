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
	type doggosOrError struct {
		doggos []model.Doggo
		err    error
	}
	doggosChannel := make(chan doggosOrError, 1)
	defer close(doggosChannel)
	go func() {
		r, e := d.repository.GetDoggos(page, limit, breedID)
		if e != nil {
			doggosChannel <- doggosOrError{doggos: nil, err: e}
		} else {
			doggosChannel <- doggosOrError{doggos: d.mapper.Map(r), err: nil}
		}
	}()
	r := <-doggosChannel
	return r.doggos, r.err
}
