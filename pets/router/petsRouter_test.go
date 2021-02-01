package router_test

import (
	router "DoggosPkg/pets/router"
	usecase "DoggosPkg/pets/usecase/mocks"
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
	mockUseCase := new(usecase.PetsUseCase)
	router.NewPetsHandler(e, mockUseCase)

	pets := util.GeneratePets(5)
	mockUseCase.On("GetPets", 1, 25, "").Return(pets, nil)

	req, err := http.NewRequest(echo.GET, "/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.PetsHandler{
		PetsUseCase: mockUseCase,
	}

	err = handler.FetchPets(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	expecations, _ := json.Marshal(pets)

	require.JSONEq(t, string(expecations), rec.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUseCase := new(usecase.PetsUseCase)
	mockUseCase.On("GetPets", 1, 25, "").Return(nil, errors.New(""))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.PetsHandler{
		PetsUseCase: mockUseCase,
	}
	err = handler.FetchPets(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}
