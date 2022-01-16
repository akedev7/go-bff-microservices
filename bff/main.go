package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/akedev7/go-bff-microservices/bff/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	timeout      = time.Second
	quote_client client.QuoteClient
)

func QuoteRegister(router *gin.RouterGroup) {
	router.GET("/:id", GetQuote)

}

func GetQuote(c *gin.Context) {

	id, err := getParam(c, "id")

	if err != nil {
		response(c, nil, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	data, err := quote_client.GetQuote(id, &ctx)
	response(c, data, err)
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
	QuoteRegister(api.Group("/quotes"))

	r.Run()
}
