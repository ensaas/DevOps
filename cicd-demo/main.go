package main


import (
	"github.com/gin-gonic/gin"
	"helloworld/pkg/controllers"
)

func main() {
	router := gin.Default()
	router.GET("", controllers.Helloworld)
	router.Run(":8080")
}