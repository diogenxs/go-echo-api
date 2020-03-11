package main

import (
	"github.com/diogenxs/go-echo-api/handler"
)

func main() {
	e := handler.NewRouter()
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
