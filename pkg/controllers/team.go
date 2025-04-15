package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var teams = []gin.H{}

func CreateTeam(c *gin.Context) {
	var team gin.H
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	teams = append(teams, team)
	c.JSON(http.StatusOK, gin.H{"team": team})
}
