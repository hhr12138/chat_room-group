package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/chat_room/service"
)

func AddUserToChat(ctx *gin.Context){
	request := new(service.AddUserRequest)
	service.AddUserToChat(request)
}
