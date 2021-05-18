package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"helloworld/pkg/services"

	
)

func Helloworld(c *gin.Context)  {
	Version, _ := services.Output()
	c.String(http.StatusOK, "Hello World CI/CD %s", Version)
}