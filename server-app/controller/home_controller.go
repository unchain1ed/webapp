package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// func CORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// クロスオリジン用にセット
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
// 		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		// w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
//         // w.Header().Set("Content-Type", "application/json")
// 	})
// }

func getTop(c *gin.Context,w http.ResponseWriter) {
	// next := http.Handler.ServeHTTP()
	// CORS()
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "home.html", gin.H{})
	// c.HTML(http.StatusOK, "home.html", nil)
}




// func postSignup(c *gin.Context) {
// 	id := c.PostForm("user_id")
// 	pw := c.PostForm("password")
// 	user, err := model.Signup(id,pw)
// 	if err != nil {
// 		c.Redirrect(301, "/signup")
// 		return
// 	}
// 	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})

// }
