package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//Home home for an api
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "API for Trollstagram")
}
