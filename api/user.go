package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"trollstagram-backend/model"
	"trollstagram-backend/opencv"

	"github.com/labstack/echo"
)

//GetUserByID get specific user by id
func GetUserByID(c echo.Context) error {
	user, err := model.GetByID(1)
	if err != nil {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

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

	filePath := fmt.Sprintf("./img/storage/%s.jpg", time.Now().Format("20060102150405"))

	model.AddFilePath(filePath)

	opencv.ProcessImage(filePath)

	return c.String(http.StatusOK, "upload success")
}

//GetImageLists get a list of all images
func GetImageLists(c echo.Context) {

}
