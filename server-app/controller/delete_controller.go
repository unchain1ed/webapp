package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/unchain1ed/server-app/model/db"
)

func getDeleteBlog(c *gin.Context) {
	// c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Content-Type", "text/plain")
	//IDをリクエストから取得
	id := c.Param("id")

	blog ,err := db.DeleteBlogInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"Deleted blog.Titile": blog.Title})
}