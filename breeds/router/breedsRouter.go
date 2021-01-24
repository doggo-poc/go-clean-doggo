package router

import (
	"DoggosPkg/breeds/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type BreedsHandler struct {
	breedsUseCase usecase.BreedsUseCase
}

func NewBreedsHandler(e *echo.Echo, dR usecase.BreedsUseCase) {
	handler := &BreedsHandler{
		breedsUseCase: dR,
	}
	e.GET("/breeds", handler.FetchBreeds)
}

func (h *BreedsHandler) FetchBreeds(c echo.Context) error {
	breeds, err := h.breedsUseCase.GetBreeds()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, breeds)
}
