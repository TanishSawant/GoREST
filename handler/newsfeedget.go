package handler

import (
	"net/http"

	"github.com/GoREST/go-rest-api-gin/newsfeed"
	"github.com/gin-gonic/gin"
)

//NewsfeedGet -
func NewsfeedGet(feed newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := feed.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
