package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/livekit/protocol/auth"
	"net/http"
	"os"
	"time"
)

func GenerateTokenHandler(c *gin.Context) {
	room := c.Query("room")
	username := c.Query("username")

	if room == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 'room' 查询参数"})
		return
	}
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 'username' 查询参数"})
		return
	}

	apiKey := os.Getenv("LIVEKIT_API_KEY")
	apiSecret := os.Getenv("LIVEKIT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器配置错误"})
		return
	}

	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.AddGrant(grant).SetIdentity(username).SetValidFor(time.Hour)

	token, err := at.ToJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
