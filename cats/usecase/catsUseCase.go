package usecase

import (
	"DoggosPkg/cats/adapter"
	"DoggosPkg/cats/usecase/model"
	"DoggosPkg/repositories/cats"
)

type CatUseCase interface {
	GetCats(page int, limit int, breedID string) ([]model.Cat, error)
}

type catUseCase struct {
	repository cats.CatRepository
	mapper     adapter.CatsMapper
}

func NewCatUseCase(catsRepo cats.CatRepository, catsMapper adapter.CatsMapper) CatUseCase {
	return &catUseCase{
		repository: catsRepo,
		mapper:     catsMapper,
	}
}

func (d *catUseCase) GetCats(page int, limit int, breedID string) ([]model.Cat, error) {
	type catsOrError struct {
		cats []model.Cat
		err  error
	}
	catsChannel := make(chan catsOrError, 1)
	defer close(catsChannel)
	go func() {
		r, e := d.repository.GetCats(page, limit, breedID)
		if e != nil {
			catsChannel <- catsOrError{cats: nil, err: e}
		} else {
			catsChannel <- catsOrError{cats: d.mapper.Map(r), err: nil}
		}
	}()
	r := <-catsChannel
	return r.cats, r.err
}
