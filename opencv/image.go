package opencv

import (
	"image/color"

	"gocv.io/x/gocv"
)

var image gocv.Mat

var (
	blue          = color.RGBA{0, 0, 255, 0}
	faceAlgorithm = "haarcascade_frontalface_default.xml"
)

//OpenImage open image.jpg
func OpenImage() {
	image = gocv.IMRead("./img/image.jpg", gocv.IMReadUnchanged)
}

//DetectFace detect a face from image.jpg
func DetectFace() {

}
