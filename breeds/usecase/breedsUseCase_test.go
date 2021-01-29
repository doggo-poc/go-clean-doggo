package usecase_test

import (
	adapterReal "DoggosPkg/breeds/adapter"
	adapter "DoggosPkg/breeds/adapter/mocks"

	usecase "DoggosPkg/breeds/usecase"
	repo "DoggosPkg/repositories/breeds/mocks"
	"DoggosPkg/util"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockRepo := new(repo.BreedsRepository)
	mockRepo.On("GetBreeds").Return(nil, errors.New(""))

	mockMapper := new(adapter.BreedsMapper)

	useCase := usecase.NewBreedsUseCase(mockRepo, mockMapper)
	res, err := useCase.GetBreeds()

	assert.NotNil(t, err)
	mockMapper.AssertNotCalled(t, "Map", res)
}

func Test_successful(t *testing.T) {
	dto := util.GenerateMockedBreedDto(3)
	mockRepo := new(repo.BreedsRepository)
	mockRepo.On("GetBreeds").Return(dto, nil)

	mapper := adapterReal.NewBreedsMapper()

	useCase := usecase.NewBreedsUseCase(mockRepo, mapper)
	res, err := useCase.GetBreeds()

	assert.Nil(t, err)
	assert.Equal(t, res, mapper.Map(dto))
}
