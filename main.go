package main

import (
	"github.com/gin-gonic/gin"
	rollingFile "github.com/lanziliang/logrus-rollingfile-hook"
	"github.com/sirupsen/logrus"
	ginLogrus "github.com/toorop/gin-logrus"
	"link-collector/controllers"
	"link-collector/middlewares"
	"link-collector/models"
	"net/http"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	hook, err := rollingFile.NewRollingFileTimeHook("./logs/link-collector.log", "2006-01-02", 1)
	if err != nil {
		logger.Panicln(err)
	}

	logger.AddHook(hook)

	r := gin.Default()
	r.Use(middlewares.TokenValidator())
	models.CreateDbConnection()
	models.SetupLogger(logger)

	r.Use(ginLogrus.Logger(logger), gin.Recovery())

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
