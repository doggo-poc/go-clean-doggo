package adapter_test

import (
	"DoggosPkg/breeds/adapter"
	dtoModel "DoggosPkg/repositories/breeds/model"
	"DoggosPkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapter.NewBreedsMapper().Map(make([]dtoModel.BreedDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	values := util.GenerateMockedBreedDto(n)
	assert.Len(t, adapter.NewBreedsMapper().Map(values), n)
}

func TestMapContent(t *testing.T) {
	n := 5
	values := util.GenerateMockedBreedDto(n)
	assert.Equal(t, adapter.NewBreedsMapper().Map(values), util.GenerateMockedBreed(n))
}
