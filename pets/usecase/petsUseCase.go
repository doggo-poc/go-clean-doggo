package usecase

import (
	"DoggosPkg/pets/adapter"
	"DoggosPkg/pets/usecase/model"
	"DoggosPkg/repositories"
	"DoggosPkg/repositories/cats"
	catsModel "DoggosPkg/repositories/cats/model"
	"math/rand"
)

type doggosOrError struct {
	doggos []repositories.DoggoDto
	err    error
}
type catsOrError struct {
	cats []catsModel.CatDto
	err  error
}

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

func getDogs(dogsChan chan<- doggosOrError, doggoRepo repositories.DoggoRepository, page int, limit int) {
	r, e := doggoRepo.GetDoggos(page, limit, "")
	if e != nil {
		dogsChan <- doggosOrError{doggos: nil, err: e}
	} else {
		dogsChan <- doggosOrError{doggos: r, err: nil}
	}
}

func getCats(catsChannel chan<- catsOrError, catsRepo cats.CatRepository, page int, limit int) {
	r, e := catsRepo.GetCats(page, limit, "")
	if e != nil {
		catsChannel <- catsOrError{cats: nil, err: e}
	} else {
		catsChannel <- catsOrError{cats: r, err: nil}
	}
}

func (d *petUseCase) GetPets(page int, limit int, breedID string) ([]model.Pet, error) {

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

	go func(dogsChan chan doggosOrError, catsChan chan catsOrError) {
		go getDogs(dogsChan, d.dogsRepo, page, limit)
		go getCats(catsChan, d.catsRepo, page, limit)

		catsResult := <-catsChan
		dogsResult := <-dogsChan
		if catsResult.err != nil && dogsResult.err != nil {
			resultChannel <- petsOrError{pets: nil, err: dogsResult.err}
		} else {
			pets := d.mapper.Map(dogsResult.doggos, catsResult.cats)
			rand.Shuffle(len(pets), func(i, j int) { pets[i], pets[j] = pets[j], pets[i] })
			resultChannel <- petsOrError{pets: pets, err: nil}
		}
	}(doggosChannel, catsChannel)

	result := <-resultChannel
	return result.pets, result.err

}
