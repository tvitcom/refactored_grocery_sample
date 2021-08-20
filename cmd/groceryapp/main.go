package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/refactored_grocery_sample/internal/routes"
	"log"
)

func main() {
	log.Print("Starting the golang grocery app with templates")
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/**/*")
	routes.Setup(r)
	r.Run()
}
