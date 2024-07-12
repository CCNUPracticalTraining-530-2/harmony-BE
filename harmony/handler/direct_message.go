package handler

import (
	"github.com/gin-gonic/gin"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
	"strconv"
)

const MessagesBatch = 10

func GetDirectMessagesHandler(c *gin.Context) {
	// 获取 URL 参数
	conversationID, err := strconv.Atoi(c.Query("conversationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Conversation ID missing or invalid"})
		return
	}

	// 查询消息列表
	var messages []model.DirectMessage
	cursor := c.Query("cursor")

	query := mysql.DB.Model(&model.DirectMessage{}).
		Where("conversation_id = ?", conversationID).
		Order("created_at desc").
		Limit(MESSAGES_BATCH)

	if cursor != "" {
		query = query.Where("id < ?", cursor)
	}

	err = query.Find(&messages).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	// 获取下一个 cursor
	var nextCursor interface{}
	if len(messages) == MESSAGES_BATCH {
		nextCursor = messages[len(messages)-1].ID
	}

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"items":      messages,
		"nextCursor": nextCursor,
	})
}
