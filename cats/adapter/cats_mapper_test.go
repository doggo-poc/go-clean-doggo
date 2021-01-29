package adapter_test

import (
	"DoggosPkg/cats/adapter"
	dtoModel "DoggosPkg/repositories/cats/model"
	"DoggosPkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapter.NewCatMapper().Map(make([]dtoModel.CatDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	values := util.GenerateMockedCatsDto(n)
	assert.Len(t, adapter.NewCatMapper().Map(values), n)
}

func TestMapContent(t *testing.T) {
	n := 5
	values := util.GenerateMockedCatsDto(n)
	assert.Equal(t, adapter.NewCatMapper().Map(values), util.GenerateCat(n))
}
