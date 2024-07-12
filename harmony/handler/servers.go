package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func ServerRemoveHandler(c *gin.Context) {
	serverID := c.Param("serverId")

	profileID := c.Param("profileId")

	if profileID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete server record
	result := mysql.DB.Where("id = ? AND profile_id = ?", serverID, profileID).Delete(&model.Server{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Server deleted successfully"})
	return
}

func ServeUpdateHandler(c *gin.Context) {
	serverID := c.Param("serverId")

	profileID := c.Param("profileId")

	if profileID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Update server record
	result := mysql.DB.Model(&model.Server{}).
		Where("id = ? AND profile_id = ?", serverID, profileID).
		Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Server updated successfully"})
	return
}

func CreateServerHandler(c *gin.Context) {
	var requestData struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageUrl"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	profileID := c.Param("profileId") // Replace with actual function to get profile ID

	if profileID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Generate UUID for invite code
	inviteCode := uuid.New().String()

	// Start a transaction
	tx := mysql.DB.Begin()

	// Create server record
	server := model.Server{
		Name:       requestData.Name,
		ImageURL:   requestData.ImageURL,
		ProfileID:  profileID,
		InviteCode: inviteCode,
	}
	if err := tx.Create(&server).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Create default channel for the server
	channel := model.Channel{
		Name:      "一般頻道",
		ProfileID: profileID,
		ServerID:  server.ID,
	}
	if err := tx.Create(&channel).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Create member with ADMIN role for the current profile
	member := model.Member{
		ProfileID: profileID,
		Role:      "ADMIN", // Assuming MemberRole.ADMIN corresponds to "ADMIN" in your context
		ServerID:  server.ID,
	}
	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusOK, server)
	return
}
