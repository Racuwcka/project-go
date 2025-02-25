package main

import (
	"log"
	httpapp "route256/cart/internal/pkg/app/http"
)

func main() {
	app := httpapp.NewApp()
	log.Fatal(app.Start())
}
