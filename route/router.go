package route

import (
	"trollstagram-backend/api"

	"github.com/labstack/echo"
)

//Init start
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)
	e.POST("/image", api.UploadImage)

	return e
}
