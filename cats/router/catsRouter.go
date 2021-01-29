package router

import (
	"DoggosPkg/cats/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CatsHandler struct {
	CatsUseCase usecase.CatUseCase
}

func NewCatsHandler(e *echo.Echo, dR usecase.CatUseCase) {
	handler := &CatsHandler{
		CatsUseCase: dR,
	}
	e.GET("/cats", handler.FetchCats)
}

func (h *CatsHandler) FetchCats(c echo.Context) error {
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

	doggos, err := h.CatsUseCase.GetCats(page, limit, breedID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, doggos)
}
