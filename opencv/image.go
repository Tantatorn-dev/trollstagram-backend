package opencv

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var (
	blue          = color.RGBA{0, 0, 255, 0}
	faceAlgorithm = "haarcascade_frontalface_default.xml"
)

//ProcessImage process an image
func ProcessImage(filePath string) {
	var img gocv.Mat
	var trollface gocv.Mat
	var classifier gocv.CascadeClassifier

	//load image
	img = gocv.IMRead("./img/image.jpg", gocv.IMReadUnchanged)
	defer img.Close()

	trollface = gocv.IMRead("./img/troll.png", gocv.IMReadUnchanged)
	defer trollface.Close()

	//load classifier
	classifier = gocv.NewCascadeClassifier()
	classifier.Load(faceAlgorithm)
	defer classifier.Close()
	rects := classifier.DetectMultiScale(img)
	fmt.Printf("found %d faces\n", len(rects))

	for _, r := range rects {
		size := r.Size()
		gocv.Resize(trollface, &trollface, image.Point{size.X * 2, size.Y * 2}, 0, 0, gocv.InterpolationDefault)

		renderImg(&img, trollface, img, image.Point{r.Min.X, r.Min.Y})
	}
	gocv.IMWrite(filePath, img)
}

func renderImg(imgdst *gocv.Mat, transImg, baseImg gocv.Mat, pt image.Point) {
	*imgdst = gocv.NewMatWithSize(baseImg.Rows(), baseImg.Cols(), gocv.MatTypeCV8UC3)
	baseImg.CopyTo(imgdst)

	planesRGBA := gocv.Split(transImg)

	dx := pt.X - transImg.Cols()/3
	dy := pt.Y - transImg.Rows()/3
	for c := 0; c < len(planesRGBA)-1; c++ {
		for i := 0; i < planesRGBA[c].Rows(); i++ {
			for j := 0; j < planesRGBA[c].Cols(); j++ {
				if planesRGBA[3].GetUCharAt(i, j) != 0 {
					py := i + dy
					px := j + dx
					if px > 0 && py > 0 && px < imgdst.Cols() && py < imgdst.Rows() {
						imgdst.SetUCharAt(py, px*3+c, planesRGBA[c].GetUCharAt(i, j))
					}
				}
			}
		}
	}
}
