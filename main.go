package main

import (
	"net/http"
	"time"

	"./database"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.GormConnect()
	defer db.Close()
	db.LogMode(true)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/create_tanaka", func(c *gin.Context) {

		user := database.People{}
		now := time.Now()
		user.Name = "tanaka"
		user.CreatedAt = now
		user.UpdatedAt = now
		c.BindJSON(&user)

		db.NewRecord(user)
		db.Create(&user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	})

	router.GET("/create_suzuki", func(c *gin.Context) {

		user := database.People{}
		now := time.Now()
		user.Name = "suzuki"
		user.CreatedAt = now
		user.UpdatedAt = now
		c.BindJSON(&user)

		db.NewRecord(user)
		db.Create(&user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run(":80")
}
