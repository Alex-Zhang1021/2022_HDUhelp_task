package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	new_password := "22031121"
	r.POST("/index", func(c *gin.Context) {

		username := c.DefaultPostForm("username", "")
		password := c.DefaultPostForm("password", "")
		u := UserInfo{
			Username: username,
			Password: password,
		}
		fmt.Printf("%v\n", u)
		fmt.Printf("%v\n", new_password)

		if username == "张颢宇" && password == new_password {
			c.HTML(200, "index.html", gin.H{
				"Name":     username,
				"Password": password,
			})
		} else {
			c.HTML(200, "login.html", nil)
		}

	})
	r.POST("/change", func(c *gin.Context) {
		c.HTML(200, "change.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		password := c.DefaultPostForm("password", "")
		new_password = password
		c.HTML(200, "login.html", nil)
	})

	r.Run()

}
