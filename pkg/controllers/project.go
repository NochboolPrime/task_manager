package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var projects = []gin.H{}

func CreateProject(c *gin.Context) {
	var project gin.H
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	projects = append(projects, project)
	c.JSON(http.StatusOK, gin.H{"project": project})
}
