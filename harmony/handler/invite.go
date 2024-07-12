package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func InviteHandler(c *gin.Context) {
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
	err := mysql.DB.Model(&model.Server{}).Where("id = ? AND profile_id = ?", serverID, profileID).Update("invite_code", uuid.NewString()).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Fetch updated server data
	var server model.Server
	err = mysql.DB.Where("id = ?", serverID).First(&server).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, server)
	return
}
