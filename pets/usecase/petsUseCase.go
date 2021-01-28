package usecase

import (
	"DoggosPkg/pets/adapter"
	"DoggosPkg/pets/usecase/model"
	"DoggosPkg/repositories"
	"DoggosPkg/repositories/cats"
	catsModel "DoggosPkg/repositories/cats/model"
	"math/rand"
)

type PetsUseCase interface {
	GetPets(page int, limit int, breedID string) ([]model.Pet, error)
}

type petUseCase struct {
	dogsRepo repositories.DoggoRepository
	catsRepo cats.CatRepository
	mapper   adapter.PetsMapper
}

func NewPetsUseCase(doggoRepo repositories.DoggoRepository, catsRepo cats.CatRepository, mapper adapter.PetsMapper) PetsUseCase {
	return &petUseCase{
		dogsRepo: doggoRepo,
		catsRepo: catsRepo,
		mapper:   mapper,
	}
}

func (d *petUseCase) GetPets(page int, limit int, breedID string) ([]model.Pet, error) {
	type doggosOrError struct {
		doggos []repositories.DoggoDto
		err    error
	}
	type catsOrError struct {
		cats []catsModel.CatDto
		err  error
	}
	type petsOrError struct {
		pets []model.Pet
		err  error
	}
	doggosChannel := make(chan doggosOrError, 1)
	catsChannel := make(chan catsOrError, 1)
	resultChannel := make(chan petsOrError, 1)

	defer close(doggosChannel)
	defer close(catsChannel)
	defer close(resultChannel)

	go func() {
		go func() {
			r, e := d.dogsRepo.GetDoggos(page, limit, "")
			if e != nil {
				doggosChannel <- doggosOrError{doggos: nil, err: e}
			} else {
				doggosChannel <- doggosOrError{doggos: r, err: nil}
			}
		}()

		go func() {
			r, e := d.catsRepo.GetCats(page, limit, "")
			if e != nil {
				catsChannel <- catsOrError{cats: nil, err: e}
			} else {
				catsChannel <- catsOrError{cats: r, err: nil}
			}
		}()
		catsResult := <-catsChannel
		dogsResult := <-doggosChannel
		if catsResult.err != nil && dogsResult.err != nil {
			resultChannel <- petsOrError{pets: nil, err: dogsResult.err}
		} else {
			pets := d.mapper.Map(dogsResult.doggos, catsResult.cats)
			rand.Shuffle(len(pets), func(i, j int) { pets[i], pets[j] = pets[j], pets[i] })
			resultChannel <- petsOrError{pets: pets, err: nil}
		}
	}()

	result := <-resultChannel
	return result.pets, result.err

}
