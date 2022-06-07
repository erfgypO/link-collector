package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"link-collector/models"
	"net/http"
)

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		anonymousPaths := []string{"/user/sign-up", "/user/sign-in"}

		if !slices.Contains(anonymousPaths, c.Request.URL.Path) {
			apiToken := c.GetHeader("X-API-TOKEN")

			if apiToken == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing header X-API-TOKEN"})
			}

			userId, err := models.GetUserIdByToken(apiToken)
			if err != nil {
				models.Logger.Error(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Set("user_id", userId)
			c.Set("token", apiToken)
		}

		c.Next()
	}
}
