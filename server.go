package main

import (
	"trollstagram-backend/opencv"
	"trollstagram-backend/route"
)

func main() {
	e := route.Init()
	opencv.OpenImage()
	e.Logger.Fatal(e.Start(":1323"))
}
