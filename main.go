package main

import (
	//"fmt"
	"github.com/GoREST/go-rest-api-gin/handler"
	"github.com/GoREST/go-rest-api-gin/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	feed := newsfeed.New()
	server.GET("/newsfeed", handler.NewsfeedGet(feed))
	server.POST("/newsfeed", handler.NewsfeedPost())
	server.Run(":8080")
}
