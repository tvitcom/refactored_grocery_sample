package main

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/web"
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/webapi"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Print("Starting the golang grocery app with templates")
	r := gin.Default()
	webapi.Route(r)

	r.LoadHTMLGlob("web/templates/**/*")
	web.Route(r)
	r.Run()
}
