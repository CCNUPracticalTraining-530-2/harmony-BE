package handler

import (
	"github.com/gin-gonic/gin"
	"harmony/dao/mysql"
	"harmony/model"
	"net/http"
)

func DeleteMemberHandler(c *gin.Context) {
	memberId := c.Param("memberId")
	serverId := c.Query("serverId")

	profileId := c.Param("profileId")

	if serverId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Server ID missing"})
		return
	}

	if memberId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member ID missing"})
		return
	}

	var server model.Server
	result := mysql.DB.Where("id = ? AND profile_id = ?", serverId, profileId).First(&server)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	if server.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	mysql.DB.Model(&server).Association("Members").Delete(&model.Member{ID: memberId})

	c.JSON(http.StatusOK, server)
}

func PatchMemberHandler(c *gin.Context) {
	memberId := c.Param("memberId")
	serverId := c.Query("serverId")

	profileId := c.Param("profileId")

	if serverId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Server ID missing"})
		return
	}

	if memberId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member ID missing"})
		return
	}

	var server model.Server
	result := mysql.DB.Where("id = ? AND profile_id = ?", serverId, profileId).First(&server)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	if server.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	var role model.MemberRole
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// 开启事务
	tx := mysql.DB.Begin()

	// 找到指定的成员
	var member model.Member
	if err := tx.First(&member, memberId).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Member not found"})
		return
	}

	// 更新角色
	member.Role = role
	if err := tx.Save(&member).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
		return
	}

	// 提交事务
	tx.Commit()

	c.JSON(http.StatusOK, server)
}
