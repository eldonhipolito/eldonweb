package web

import (
	"net/http"
	"strconv"

	"github.com/eldonhipolito/eldonweb/models"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	render(c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		}, "index.html",
	)
}

func getArticle(c *gin.Context) {

	v, err := strconv.Atoi(c.Param("article_id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	article, err := models.GetArticleByID(v)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	render(c,
		gin.H{
			"title":   article.Title,
			"payload": article,
		}, "article.html")

}
