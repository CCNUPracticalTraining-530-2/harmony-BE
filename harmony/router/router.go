package router

import (
	"github.com/gin-gonic/gin"
	"harmony/handler"
	"harmony/router/middleware"
)

func InitRouter() *gin.Engine {
	e := gin.Default()

	e.Use(middleware.Cors)
	e.DELETE("/api/chanel/servers/:serverId/channels/:channelId/profiles/:profileId", handler.DeleteChannel)
	e.PATCH("/api/chanel/server/:serverId/channel/:channelId", handler.UpdateChannel)
	e.POST("/api/chanel/server/:serverId/channels/:channelId", handler.CreateChannelHandler)
	e.GET("/api/direct-messages", handler.GetDirectMessagesHandler)
	e.GET("api/token", handler.GenerateTokenHandler)
	e.DELETE("/api/member/server/:serverId/member/:memberId", handler.DeleteMemberHandler)
	e.PATCH("/api/member/server/:serverId/member/:memberId", handler.PatchMemberHandler)
	e.GET("/api/messages", handler.GetMessagesHandler)
	e.PATCH("/api/invite/servers/:serverId/profiles/:profileId", handler.InviteHandler)
	e.PATCH("/api/leave/servers/:serverId/profiles/:profileId", handler.LeaveHandler)
	e.DELETE("/api/server/servers/:serverId/profileId/:profileId", handler.ServerRemoveHandler)
	e.PATCH("/api/server/servers/:serverId/profileId/:profileId", handler.ServeUpdateHandler)
	e.POST("/api/server/servers/profileId/:profileId", handler.CreateServerHandler)
	e.POST("/api/upload/serveImg", handler.UploadServerImageHandler)
	e.POST("/api/upload/messageFile", handler.UploadMessageFileHandler)

	return e
}
