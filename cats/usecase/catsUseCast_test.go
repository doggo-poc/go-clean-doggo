package usecase_test

import (
	adapterReal "DoggosPkg/cats/adapter"
	adapter "DoggosPkg/cats/adapter/mocks"

	usecase "DoggosPkg/cats/usecase"
	repo "DoggosPkg/repositories/cats/mocks"
	"DoggosPkg/util"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockRepo := new(repo.CatRepository)
	mockRepo.On("GetCats", 1, 25, "").Return(nil, errors.New(""))

	mockMapper := new(adapter.CatsMapper)

	useCase := usecase.NewCatUseCase(mockRepo, mockMapper)
	res, err := useCase.GetCats(1, 25, "")

	assert.NotNil(t, err)
	mockMapper.AssertNotCalled(t, "Map", res)
}

func Test_successful(t *testing.T) {
	limit := 25
	dto := util.GenerateMockedCatsDto(limit)
	mockRepo := new(repo.CatRepository)
	mockRepo.On("GetCats", 1, limit, "").Return(dto, nil)

	mapper := adapterReal.NewCatMapper()

	useCase := usecase.NewCatUseCase(mockRepo, mapper)
	res, err := useCase.GetCats(1, limit, "")

	assert.Nil(t, err)
	assert.Equal(t, res, mapper.Map(dto))
}
