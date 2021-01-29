package adapter_test

import (
	"DoggosPkg/pets/adapter"
	doggosDtoModel "DoggosPkg/repositories"
	catsDtoModel "DoggosPkg/repositories/cats/model"
	"DoggosPkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapter.NewPetMapper().Map(make([]doggosDtoModel.DoggoDto, 0), make([]catsDtoModel.CatDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	cats := util.GenerateMockedCatsDto(n)
	dogs := util.GenerateMockedDogsDto(n)
	assert.Len(t, adapter.NewPetMapper().Map(dogs, cats), n+n)
}

func TestMapContent(t *testing.T) {
	n := 1
	cats := util.GenerateMockedCatsDto(n)
	dogs := util.GenerateMockedDogsDto(n)

	assert.Equal(t, adapter.NewPetMapper().Map(dogs, cats), util.GeneratePetsFromDoggsCats(dogs, cats))
}
