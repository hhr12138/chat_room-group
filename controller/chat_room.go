package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/chat_room/service"
	"github.com/hhr12138/chat_room/utils"
	"github.com/hhr12138/door/entity"
	"net/http"
)

func CreateChatRoom(ctx *gin.Context) {
	request := new(service.CreateChatRoomRequest)
	err := ctx.ShouldBind(request)
	request.UserId = utils.GetUser(ctx).Id
	if err != nil {
		ctx.JSON(http.StatusFound, err.Error())
		return
	}
	room, err := service.CreateChatRoom(request)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
	} else{
		ctx.JSON(http.StatusOK,room)
	}
}

func DeleteChatRoom(ctx *gin.Context){
	request := new(service.DeleteChatRoomRequest)
	err := ctx.ShouldBind(request)
	if err != nil{
		ctx.JSON(http.StatusFound,err)
		return
	}
	room, err := service.DeleteChatRoom(request)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
 	} else{
		 ctx.JSON(http.StatusOK,room)
	}
}

func GetAllUserId(ctx *gin.Context) {
	chatId := ctx.GetInt64("chat_id")
	if chatId == 0 {
		ctx.JSON(http.StatusFound, "chat_id can not nil")
		return
	}
	userIds, err := service.GetAllUserId(chatId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, userIds)
	}
}

func GetGroupIdsByUserId(ctx *gin.Context){
	u,_ := ctx.Get("user")
	user := u.(*entity.User)
	uid := user.Id
	reply, err := service.GetGroupIdsByUserId(uid)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
	} else {
		ctx.JSON(http.StatusOK, reply)
	}
}

func GetGroupsByUserId(ctx *gin.Context){
	uid := utils.GetUser(ctx).Id
	groupIds,err := service.GetGroupIdsByUserId(uid)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	groups,err := service.GetGroupsByUserId(groupIds)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,groups)
}
