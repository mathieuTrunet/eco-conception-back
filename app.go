package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Data []struct {
		Text string `json:"text" binding:"required"`
	} `json:"data" binding:"required"`
}

func main() {

	application := gin.Default()

	application.POST("/api/page/download", func(response *gin.Context) {
		var requestBody RequestBody

		if error := response.ShouldBindJSON(&requestBody); error != nil {
			response.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		template, error := template.ParseFiles("template.html")
		if error != nil {
			response.JSON(http.StatusInternalServerError, gin.H{"error": "template not found"})
			return
		}

		htmlBuffer := new(strings.Builder)

		error = template.Execute(htmlBuffer, requestBody)
		if error != nil {
			response.JSON(http.StatusInternalServerError, gin.H{"error": "parsing error"})
			return
		}

		response.Header("Content-Type", "text/html")
		response.Header("Content-Disposition", "attachment; filename=builded_page.html")

		response.String(http.StatusOK, htmlBuffer.String())
	})
	application.Run(":8080")

}
