package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {
	router.GET("/:id", GetUserDetails)

}

func GetUserDetails(c *gin.Context) {
	id, err := getParam(c, "id")
	response(c, id, err)
	return
}

func getParam(c *gin.Context, param string) (string, error) {
	p := c.Param(param)
	if len(p) == 0 {
		return "", errors.New("invalid parameter: " + p)
	}
	return p, nil
}

func response(c *gin.Context, data interface{}, err error) {
	statusCode := http.StatusOK
	var errorMessage string
	if err != nil {
		log.Println("Server Error Occured:", err)
		errorMessage = strings.Title(err.Error())
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, gin.H{"data": data, "error": errorMessage})
}

func main() {
	log.Println("Bff Service")

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	UserRegister(api.Group("/users"))

	r.Run()
}
