package repositories_test

import (
	"DoggosPkg/repositories"

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
	repo := repositories.NewDoggoRepository(mockClient)
	limit := 25
	url := repositories.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	assert.NoError(t, error)

	mockClient.On("Execute", req).Return(nil, errors.New("error"))

	_, err := repo.GetDoggos(1, limit, "")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}

func Test_EmptyResponseEmptyDataSet(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := repositories.NewDoggoRepository(mockClient)
	limit := 25
	url := repositories.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	assert.NoError(t, error)
	emptyResponse, error := json.Marshal(make([]repositories.DoggoDto, 0))
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(emptyResponse)))}, nil)

	resp, err := repo.GetDoggos(1, limit, "")

	assert.Nil(t, err)
	assert.Empty(t, resp)
}

func Test_Response(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := repositories.NewDoggoRepository(mockClient)
	limit := 25
	url := repositories.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	data := util.GenerateMockedDogsDto(limit)
	dataJson, error := json.Marshal(data)
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(dataJson)))}, nil)

	resp, err := repo.GetDoggos(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, data, resp)
}
