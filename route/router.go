package route

import (
	"trollstagram-backend/api"

	"github.com/labstack/echo"
)

//Init start
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)
	e.GET("/user", api.GetUserByID)
	e.POST("/user/image", api.UploadImage)

	return e
}
