package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			session.Set("token", true)
			session.Save()
			c.JSON(200, gin.H{"login": false})
		} else {
			c.JSON(200, gin.H{"login": true})
		}
	})
	r.Run(":8000")
}
