package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = []gin.H{}

func CreateTask(c *gin.Context) {
	var task gin.H
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}
