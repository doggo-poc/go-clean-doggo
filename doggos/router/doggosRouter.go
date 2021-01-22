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
	limitS := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitS)

	pageS := c.QueryParam("page")
	page, _ := strconv.Atoi(pageS)

	doggos, err := h.doggoUseCase.GetDoggos(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	//c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, doggos)
}
