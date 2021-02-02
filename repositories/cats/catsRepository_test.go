package cats_test

import (
	"DoggosPkg/repositories/cats"
	"DoggosPkg/repositories/cats/model"

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
	repo := cats.NewCatRepository(mockClient)
	limit := 25
	url := cats.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	assert.NoError(t, error)

	mockClient.On("Execute", req).Return(nil, errors.New("error"))

	_, err := repo.GetCats(1, limit, "")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}

func Test_EmptyResponseEmptyDataSet(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := cats.NewCatRepository(mockClient)
	limit := 25
	url := cats.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	assert.NoError(t, error)
	emptyResponse, error := json.Marshal(make([]model.CatDto, 0))
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(emptyResponse)))}, nil)

	resp, err := repo.GetCats(1, limit, "")

	assert.Nil(t, err)
	assert.Empty(t, resp)
}

func Test_Response(t *testing.T) {
	mockClient := new(mocks.HttpClient)
	repo := cats.NewCatRepository(mockClient)
	limit := 25
	url := cats.NewUrl(1, limit, "")
	req, error := http.NewRequest("GET", url, nil)

	data := util.GenerateMockedCatsDto(limit)
	dataJson, error := json.Marshal(data)
	assert.NoError(t, error)
	mockClient.On("Execute", req).Return(&http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte(dataJson)))}, nil)

	resp, err := repo.GetCats(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, data, resp)
}
