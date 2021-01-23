package router

import (
	"DoggosPkg/doggos/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

// DoggosHandler  represent the httphandler for Doggos
type DoggosHandler struct {
	doggoUseCase usecase.DoggoUseCase
}

func NewDoggosHandler(e *echo.Echo, dR usecase.DoggoUseCase) {
	handler := &DoggosHandler{
		doggoUseCase: dR,
	}
	e.GET("/", handler.FetchDoggos)
}

func (h *DoggosHandler) FetchDoggos(c echo.Context) error {
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

	doggos, err := h.doggoUseCase.GetDoggos(page, limit, breedID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, doggos)
}
