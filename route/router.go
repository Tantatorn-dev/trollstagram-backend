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
	e.GET("/user/posts", api.CountPosts)
	e.GET("/user/image", api.GetImageLists)
	e.GET("/user/image/:id", api.GetImage)
	e.POST("/user/image", api.UploadImage)

	return e
}
