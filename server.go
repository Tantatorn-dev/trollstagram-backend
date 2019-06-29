package main

import (
	"trollstagram-backend/db"
	"trollstagram-backend/route"
)

func main() {
	db.Init()
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
