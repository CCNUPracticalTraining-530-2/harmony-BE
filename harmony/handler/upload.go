package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func UploadServerImageHandler(c *gin.Context) {
	profileID := c.MustGet("profileID").(string)

	// Generate UUID for image URL
	imageURL := uuid.New().String()

	// Start a transaction
	tx := mysql.DB.Begin()

	// Create serverImage record
	serverImage := model.ServerImages{
		ImageURL:  imageURL,
		ProfileID: profileID,
	}
	if err := tx.Create(&serverImage).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload server image"})
		return
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"imageURL": imageURL})
	return
}

func UploadMessageFileHandler(c *gin.Context) {
	profileID := c.MustGet("profileID").(string)

	// Simulating file upload logic, replace with actual implementation
	filePath := "/path/to/messageFile"

	// Start a transaction
	tx := mysql.DB.Begin()

	// Create messageFile record
	messageFile := model.MessageFile{
		FilePath:  filePath,
		ProfileID: profileID,
	}
	if err := tx.Create(&messageFile).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload message file"})
		return
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"filePath": filePath})
}
