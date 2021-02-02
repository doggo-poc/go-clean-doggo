package adapters_test

import (
	"DoggosPkg/doggos/adapters"
	dtoModel "DoggosPkg/repositories"
	"DoggosPkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapters.NewDoggoMapper().Map(make([]dtoModel.DoggoDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	values := util.GenerateMockedDogsDto(n)
	assert.Len(t, adapters.NewDoggoMapper().Map(values), n)
}

func TestMapContent(t *testing.T) {
	n := 5
	values := util.GenerateMockedDogsDto(n)
	assert.Equal(t, adapters.NewDoggoMapper().Map(values), util.GenerateDoggos(n))
}
