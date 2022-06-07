package controllers

import (
	"github.com/gin-gonic/gin"
	"link-collector/models"
	"net/http"
)

func SignUp(c *gin.Context) {
	var login models.UserLoginDto

	err := c.ShouldBindJSON(&login)
	if err != nil {
		models.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = models.CreateUser(login)
	if err != nil {
		models.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, _ := models.CreateAccessToken(login)
	c.JSON(http.StatusOK, accessToken)
}

func SignIn(c *gin.Context) {
	var login models.UserLoginDto

	err := c.ShouldBindJSON(&login)
	if err != nil {
		models.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := models.CreateAccessToken(login)
	if err != nil {
		models.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, accessToken)
}

func LogOut(c *gin.Context) {
	accessToken := c.MustGet("token").(string)

	models.DeleteAccessToken(accessToken)
	c.Status(http.StatusOK)
}
