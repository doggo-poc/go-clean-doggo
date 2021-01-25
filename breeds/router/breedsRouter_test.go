package router_test

import (
	router "DoggosPkg/breeds/router"
	usecase "DoggosPkg/breeds/usecase/mocks"
	"DoggosPkg/util"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyMap(t *testing.T) {
	e := echo.New()
	mockUseCase := new(usecase.BreedsUseCase)
	router.NewBreedsHandler(e, mockUseCase)

	breeds := util.GenerateMockedBreed(5)
	mockUseCase.On("GetBreeds").Return(breeds, nil)

	req, err := http.NewRequest(echo.GET, "/breeds", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.BreedsHandler{
		BreedsUseCase: mockUseCase,
	}

	err = handler.FetchBreeds(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	expecations, _ := json.Marshal(breeds)

	require.JSONEq(t, string(expecations), rec.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUseCase := new(usecase.BreedsUseCase)
	mockUseCase.On("GetBreeds").Return(nil, errors.New(""))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/breeds", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.BreedsHandler{
		BreedsUseCase: mockUseCase,
	}
	err = handler.FetchBreeds(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}
