package main

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Print("Starting the golang grocery app with templates")
	r := gin.Default()
	r.LoadHTMLGlob("./ui/templates/**/*")
	routes.Route(r)
	r.Run()
}
