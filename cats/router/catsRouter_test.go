package router_test

import (
	router "DoggosPkg/cats/router"
	usecase "DoggosPkg/cats/usecase/mocks"
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
	mockUseCase := new(usecase.CatUseCase)
	router.NewCatsHandler(e, mockUseCase)

	cats := util.GenerateCat(5)
	mockUseCase.On("GetCats", 1, 25, "").Return(cats, nil)

	req, err := http.NewRequest(echo.GET, "/cats", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.CatsHandler{
		CatsUseCase: mockUseCase,
	}

	err = handler.FetchCats(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	expecations, _ := json.Marshal(cats)

	require.JSONEq(t, string(expecations), rec.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUseCase := new(usecase.CatUseCase)
	mockUseCase.On("GetCats", 1, 25, "").Return(nil, errors.New(""))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cats", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.CatsHandler{
		CatsUseCase: mockUseCase,
	}
	err = handler.FetchCats(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}
