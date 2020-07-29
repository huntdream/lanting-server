package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

//GetArticle get article by id
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	article, err := service.GetArticleByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
	}

	c.JSON(http.StatusOK, article)
}
