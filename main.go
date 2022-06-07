package main

import (
	"github.com/gin-gonic/gin"
	"link-collector/controllers"
	"link-collector/middlewares"
	"link-collector/models"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.TokenValidator())
	models.CreateDbConnection()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/user/sign-up", controllers.SignUp)
	r.POST("/user/sign-in", controllers.SignIn)
	r.DELETE("/user/logout", controllers.LogOut)

	r.GET("/link/all", controllers.GetLinks)
	r.POST("/link", controllers.AddLink)
	r.PUT("/link/:id", controllers.UpdateLink)
	r.DELETE("/link/:id", controllers.DeleteLink)

	r.Run(":8420")
}
