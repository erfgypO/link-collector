package controllers

import (
	"github.com/gin-gonic/gin"
	"link-collector/models"
	"net/http"
	"strconv"
)

func AddLink(c *gin.Context) {
	var linkDto models.LinkDto

	err := c.ShouldBindJSON(&linkDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("user_id").(uint)

	link := models.Link{
		Name:   linkDto.Name,
		Url:    linkDto.Url,
		UserID: userId,
	}

	models.CreateLink(&link)

	c.JSON(http.StatusOK, link)
}

func GetLinks(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)

	links := models.GetLinks(userId)

	c.JSON(http.StatusOK, links)
}

func DeleteLink(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)
	linkId, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link id"})
		return
	}

	models.DeleteLink(uint(linkId), userId)
	c.Status(http.StatusOK)
}

func UpdateLink(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)
	linkId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link id"})
		return
	}

	var linkDto models.LinkDto

	err = c.ShouldBindJSON(&linkDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link := models.Link{Id: uint(linkId)}

	models.DB.Where("user_id = ?", userId).Find(&link)

	link.Url = linkDto.Url
	link.Name = linkDto.Name

	models.DB.Save(&link)
}
