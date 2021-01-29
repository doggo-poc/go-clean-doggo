package router_test

import (
	router "DoggosPkg/doggos/router"
	usecase "DoggosPkg/doggos/usecase/mocks"
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
	mockUseCase := new(usecase.DoggoUseCase)
	router.NewDoggosHandler(e, mockUseCase)

	limit := 25
	doggos := util.GenerateDoggos(limit)
	mockUseCase.On("GetDoggos", 1, limit, "").Return(doggos, nil)

	req, err := http.NewRequest(echo.GET, "/doggos", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.DoggosHandler{
		DoggoUseCase: mockUseCase,
	}

	err = handler.FetchDoggos(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	expecations, _ := json.Marshal(doggos)

	require.JSONEq(t, string(expecations), rec.Body.String())

	mockUseCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUseCase := new(usecase.DoggoUseCase)
	mockUseCase.On("GetDoggos", 1, 25, "").Return(nil, errors.New(""))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/doggos", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := router.DoggosHandler{
		DoggoUseCase: mockUseCase,
	}
	err = handler.FetchDoggos(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUseCase.AssertExpectations(t)
}
