package main

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/webapi"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Print("Starting the grocery app")
	r := gin.Default()
	webapi.Route(r)
	r.Run()
}
