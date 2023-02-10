package chat_room

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/chat_room-group/controller"
	"github.com/hhr12138/chat_room-group/object"
	"github.com/hhr12138/door/filter"
	"github.com/jmoiron/sqlx"
)

func Register(engine *gin.Engine, db *sqlx.DB){
	object.RegisterDb(db)
	needLoginGroup := engine.Group("/login", filter.LoginFilter)
	needLoginGroup.POST("/chat_room/create",controller.CreateChatRoom)
	needLoginGroup.POST("/chat_room/delete",controller.DeleteChatRoom)
	needLoginGroup.POST("/chat_room/add_user",controller.AddUserToChat)
	needLoginGroup.GET("/chat_room/get_all_user_id",controller.GetAllUserId)
	needLoginGroup.GET("/chat_room/groupIds/byUid",controller.GetGroupIdsByUserId)
	needLoginGroup.GET("/chat_room/groups/byUid",controller.GetGroupsByUserId)
}