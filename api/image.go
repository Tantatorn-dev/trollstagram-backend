package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

//UploadImage upload an image request
func UploadImage(c echo.Context) error {

	//Read file

	file, err := c.FormFile("imageFile")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Sprintln(err)
		return err
	}
	defer src.Close()

	//Destination
	dst, err := os.Create("./img/image.jpg")
	if err != nil {
		return err
	}
	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "upload success")
}
