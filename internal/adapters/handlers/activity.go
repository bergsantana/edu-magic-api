package handlers

import (
	"strconv"

	"github.com/bergsantana/edu-magic-api/internal/core/domain"
	"github.com/bergsantana/edu-magic-api/internal/core/services"
	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	activityService services.ActivityService
}

func NewActivityHandler(activityService services.ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var activity domain.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.activityService.CreateActivity(&activity); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create activity"})
		return
	}
	c.JSON(201, gin.H{"message": "Activity created successfully"})
}

func (h *ActivityHandler) GetActivitiesByUser(c *gin.Context) {
	ownerIDStr := c.Param("ownerID")
	ownerID, err := strconv.ParseInt(ownerIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid owner ID"})
		return
	}
	activities, err := h.activityService.GetActivitiesByUser(ownerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve activities"})
		return
	}
	c.JSON(200, activities)
}
