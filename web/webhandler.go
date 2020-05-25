package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.GET("/articles/:article_id", getArticle)

	portNumber := os.Getenv("PORT")

	if len(portNumber) == 0 {
		portNumber = "8080"
	}
	router.Run(fmt.Sprintf(":%v", portNumber))

}

func render(c *gin.Context, d gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {

	case "application/json":
		c.JSON(http.StatusOK, d)
	case "application/xml":
		c.XML(http.StatusOK, d)
	default:
		c.HTML(http.StatusOK, templateName, d)
	}

}
