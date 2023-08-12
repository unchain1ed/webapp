package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/db"
)

func getDeleteBlog(c *gin.Context) {
	//IDをリクエストから取得
	id := c.Param("id")

	blog ,err := db.DeleteBlogInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	log.Println("Deleted blog.Titile", blog.Title);
	
	c.JSON(http.StatusOK, gin.H{"Deleted blog.Titile": blog.Title})
}