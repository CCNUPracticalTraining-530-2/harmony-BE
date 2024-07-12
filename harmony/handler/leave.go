package handler

import (
	"github.com/gin-gonic/gin"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func LeaveHandler(c *gin.Context) {
	serverID := c.Param("serverId")

	profileID := c.Param("profileId")

	if profileID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if serverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Server ID Missing"})
		return
	}

	// Update server with new invite code
	var server model.Server
	result := mysql.DB.Where("id = ? AND profile_id != ?", serverID, profileID).First(&server)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	// Fetch updated server data
	var server1 model.Server
	err := mysql.DB.Where("id = ?", serverID).First(&server1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, server)
	return
}
