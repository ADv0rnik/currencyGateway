package main

import (
	"currencyGateway/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := impl.ReturnRouter()

	queryParams := map[string]string{
		"baseUrl": "https://belarusbank.by/api/kursExchange",
		"city":    "Гомель",
	}
	response, _ := impl.SendRequest(queryParams)
	r.GET("/currency", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, response)
	})
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
