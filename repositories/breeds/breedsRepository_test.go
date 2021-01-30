package breeds_test

import (
	"DoggosPkg/repositories/breeds"
	"DoggosPkg/repositories/breeds/model"
	mocks "DoggosPkg/repositories/mocks"
	"DoggosPkg/util"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := breeds.NewBreedsRepository(mockClient)
	req, error := http.NewRequest("GET", "https://api.thedogapi.com/v1/breeds", nil)

	assert.NoError(t, error)

	mockClient.On("Execute", req).Return(nil, errors.New("error"))

	_, err := repo.GetBreeds()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}

func Test_EmptyResponseEmptyDataSet(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := breeds.NewBreedsRepository(mockClient)
	req, error := http.NewRequest("GET", "https://api.thedogapi.com/v1/breeds", nil)

	assert.NoError(t, error)
	emptyResponse, error := json.Marshal(make([]model.BreedDto, 0))
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(emptyResponse)))}, nil)

	resp, err := repo.GetBreeds()

	assert.Nil(t, err)
	assert.Empty(t, resp)
}

func Test_Response(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := breeds.NewBreedsRepository(mockClient)
	req, error := http.NewRequest("GET", "https://api.thedogapi.com/v1/breeds", nil)

	assert.NoError(t, error)
	data := util.GenerateMockedBreedDto(2)
	dataJson, error := json.Marshal(data)
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(dataJson)))}, nil)

	resp, err := repo.GetBreeds()

	assert.Nil(t, err)
	assert.ElementsMatch(t, data, resp)
}
