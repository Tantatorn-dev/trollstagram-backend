package main

import (
	"trollstagram-backend/route"
)

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
