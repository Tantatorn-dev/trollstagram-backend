package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

	filePath := fmt.Sprintf("./img/storage/%d.jpg", model.CountPosts()+1)

	model.AddFilePath(filePath)

	opencv.ProcessImage(filePath)

	return c.String(http.StatusOK, "upload success")
}

//CountPosts count all posts
func CountPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, model.CountPosts())
}

//GetImageLists get a list of all images
func GetImageLists(c echo.Context) error {
	return c.JSON(http.StatusOK, model.GetFilePaths())
}

//GetImage get image file
func GetImage(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	filepath := fmt.Sprintf("./img/storage/%d.jpg", id)
	return c.File(filepath)
}
