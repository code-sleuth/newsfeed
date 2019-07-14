package handler

import (
	"net/http"

	"github.com/code-sleuth/newsfeed/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := feed.GetAll()
		c.JSON(http.StatusOK, res)
	}
}
