package router

import (
	"DoggosPkg/pets/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PetsHandler struct {
	petsUseCase usecase.PetsUseCase
}

func NewPetsHandler(e *echo.Echo, dR usecase.PetsUseCase) {
	handler := &PetsHandler{
		petsUseCase: dR,
	}
	e.GET("/", handler.FetchPets)
}

func (h *PetsHandler) FetchPets(c echo.Context) error {
	queryLimit := c.QueryParam("limit")
	var limit int
	if queryLimit == "" {
		limit = 25
	} else {
		limit, _ = strconv.Atoi(queryLimit)
	}

	var queryPageSize = c.QueryParam("page")
	var page int
	if queryPageSize == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(queryPageSize)
	}

	breedID := c.QueryParam("breed_id")

	doggos, err := h.petsUseCase.GetPets(page, limit, breedID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, doggos)
}
