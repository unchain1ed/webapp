package controller

// import (
// 	"github.com/unchain1ed/gin-app/model"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func getSignup(c *gin.Context) {
// 	c.HTML(http.StatusOK, "signup.html", nil)
// }

// func postSignup(c *gin.Context) {
// 	id := c.PostForm("user_id")
// 	pw := c.PostForm("password")
// 	user, err := model.Signup(id, pw)
// 	if err != nil {
// 		c.Redirect(301, "/signup")
// 		return
// 	}
// 	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
// }

// func getLogin(c *gin.Context) {
// 	c.HTML(http.StatusOK, "login.html", nil)
// }

// func postLogin(c *gin.Context) {
// 	id := c.PostForm("user_id")
// 	pw := c.PostForm("password")

// 	user, err := model.Login(id, pw)
// 	if err != nil {
// 		c.Redirect(301, "/login")
// 		return
// 	}
// 	c.HTML(http.StatusOK, "top.html", gin.H{"user": user})
// }
