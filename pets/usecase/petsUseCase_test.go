package usecase_test

import (
	adapterReal "DoggosPkg/pets/adapter"
	adapter "DoggosPkg/pets/adapter/mocks"

	"DoggosPkg/util"

	usecase "DoggosPkg/pets/usecase"
	catsRepo "DoggosPkg/repositories/cats/mocks"
	dogsRepo "DoggosPkg/repositories/mocks"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockCatsRepo.On("GetCats", 1, 25, "").Return(nil, errors.New(""))
	mockDogsRepo.On("GetDoggos", 1, 25, "").Return(nil, errors.New(""))

	mockMapper := new(adapter.PetsMapper)

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mockMapper)
	res, err := useCase.GetPets(1, 25, "")

	assert.NotNil(t, err)
	mockMapper.AssertNotCalled(t, "Map", res)
}

func Test_successful(t *testing.T) {
	limit := 25
	catsDto := util.GenerateMockedCatsDto(limit)
	dogsDto := util.GenerateMockedDogsDto(limit)
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockCatsRepo.On("GetCats", 1, limit, "").Return(catsDto, nil)
	mockDogsRepo.On("GetDoggos", 1, limit, "").Return(dogsDto, nil)

	mapper := adapterReal.NewPetMapper()

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mapper)
	res, err := useCase.GetPets(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, res, mapper.Map(dogsDto, catsDto))
}

func Test_successfulDogsCatsError(t *testing.T) {
	limit := 25

	dogsDto := util.GenerateMockedDogsDto(limit)
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockCatsRepo.On("GetCats", 1, limit, "").Return(nil, errors.New(""))
	mockDogsRepo.On("GetDoggos", 1, limit, "").Return(dogsDto, nil)

	mapper := adapterReal.NewPetMapper()

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mapper)
	res, err := useCase.GetPets(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, res, mapper.Map(dogsDto, nil))
}

func Test_successfulDogsCatsEmpty(t *testing.T) {
	limit := 25

	catsDto := util.GenerateMockedCatsDto(0)
	assert.Empty(t, len(catsDto))
	dogsDto := util.GenerateMockedDogsDto(limit)
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockCatsRepo.On("GetCats", 1, limit, "").Return(catsDto, nil)
	mockDogsRepo.On("GetDoggos", 1, limit, "").Return(dogsDto, nil)

	mapper := adapterReal.NewPetMapper()

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mapper)
	res, err := useCase.GetPets(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, res, mapper.Map(dogsDto, catsDto))
}

func Test_successfulCatsDogsError(t *testing.T) {
	limit := 25

	catsDto := util.GenerateMockedCatsDto(limit)
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockDogsRepo.On("GetDoggos", 1, limit, "").Return(nil, errors.New(""))
	mockCatsRepo.On("GetCats", 1, limit, "").Return(catsDto, nil)

	mapper := adapterReal.NewPetMapper()

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mapper)
	res, err := useCase.GetPets(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, res, mapper.Map(nil, catsDto))
}

func Test_successfulCatsDogsEmpty(t *testing.T) {
	limit := 25

	dogsDto := util.GenerateMockedDogsDto(0)
	assert.Empty(t, len(dogsDto))
	catsDto := util.GenerateMockedCatsDto(limit)
	mockCatsRepo := new(catsRepo.CatRepository)
	mockDogsRepo := new(dogsRepo.DoggoRepository)
	mockCatsRepo.On("GetCats", 1, limit, "").Return(catsDto, nil)
	mockDogsRepo.On("GetDoggos", 1, limit, "").Return(dogsDto, nil)

	mapper := adapterReal.NewPetMapper()

	useCase := usecase.NewPetsUseCase(mockDogsRepo, mockCatsRepo, mapper)
	res, err := useCase.GetPets(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, res, mapper.Map(dogsDto, catsDto))
}
