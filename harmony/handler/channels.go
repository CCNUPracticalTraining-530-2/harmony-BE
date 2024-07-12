package handler

import (
	"github.com/gin-gonic/gin"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func CreateChannelHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的JSON数据"})
		return
	}

	serverID := c.Query("serverId")
	if serverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少服务器ID"})
		return
	}

	ProfileID := c.Query("profileId")
	if ProfileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少服务器ID"})
		return
	}

	if req.Name == "一般" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "名称不能为'一般頻道'"})
		return
	}

	var server model.Server
	if err := mysql.DB.Preload("Channels").First(&server, "id = ? AND members.profile_id = ? AND members.role IN (?, ?)",
		serverID, ProfileID, model.ADMIN, model.MODERATOR).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权修改服务器"})
		return
	}

	channel := model.Channel{
		ProfileID: ProfileID,
		Name:      req.Name,
		Type:      model.ChannelType(req.Type),
	}

	// 使用事务确保原子性操作
	tx := mysql.DB.Begin()
	if err := tx.Model(&server).Association("Channels").Append(&channel); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建频道失败"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, server)
}

func DeleteChannel(c *gin.Context) {
	serverId := c.Param("serverId")
	channelId := c.Param("channelId")
	profileId := c.Param("profileId")
	var server model.Server
	if err := mysql.DB.Preload("Members", "profile_id = ? AND role IN (?, ?)", profileId, model.ADMIN, model.MODERATOR).
		First(&server, "id = ?", serverId).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var channel model.Channel
	if err := mysql.DB.First(&channel, "id = ? AND server_id = ? AND name != ?", channelId, serverId, "一般頻道").Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Channel ID missing or invalid"})
		return
	}

	if err := mysql.DB.Delete(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	c.JSON(http.StatusOK, server)
}

func UpdateChannel(c *gin.Context) {
	serverId := c.Param("serverId")
	channelId := c.Param("channelId")
	profileId := c.Param("profileId")

	var server model.Server
	if err := mysql.DB.Preload("Members", "profile_id = ? AND role IN (?, ?)", profileId, model.ADMIN, model.MODERATOR).
		First(&server, "id = ?", serverId).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var channel model.Channel
	if err := mysql.DB.First(&channel, "id = ? AND server_id = ? AND name != ?", channelId, serverId, "一般頻道").Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Channel ID missing or invalid"})
		return
	}

	var json struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if json.Name == "一般頻道" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be '一般頻道'"})
		return
	}

	channel.Name = json.Name
	channel.Type = model.ChannelType(json.Type)

	if err := mysql.DB.Save(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	c.JSON(http.StatusOK, server)
}
