package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

const MESSAGES_BATCH = 10

func GetMessagesHandler(c *gin.Context) {
	var messages []model.Message
	channelID := c.Query("channelId")
	cursor := c.Query("cursor")

	if channelID == "" {
		c.JSON(http.StatusBadRequest, errors.New("Channel ID missing"))
		return

	}

	query := mysql.DB.Model(&model.Message{}).Where("channel_id = ?", channelID).Order("created_at desc").Limit(MESSAGES_BATCH)

	if cursor != "" {
		query = query.Where("id < ?", cursor)
	}

	err := query.Find(&messages).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	nextCursor := ""
	if len(messages) == MESSAGES_BATCH {
		nextCursor = messages[MESSAGES_BATCH-1].ID
	}

	c.JSON(http.StatusOK, gin.H{
		"items":      messages,
		"nextCursor": nextCursor,
	})

	return
}
