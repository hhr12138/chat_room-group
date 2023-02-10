package service

import (
	"github.com/hhr12138/chat_room/mapper"
	"github.com/hhr12138/chat_room/object"
)

type AddUserRequest struct {
	object.BasicItem
	UserId     int64 `from:"user_id",binding:"required"`
	UserRole   int64 `from:"user_role",binding:"required"`
	ChatId     int64 `from:"chat_id",binding:"required"`
	OperatorId int64 `from:"operator_id",binding:"required"`
}

type AddUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func AddUserToChat(request *AddUserRequest) (*AddUserResponse, error) {
	reply := new(AddUserResponse)
	chatId := request.ChatId
	needRole, err := mapper.GetAddUserRoleById(chatId)
	if err != nil{
		return reply,err
	}
	if request.UserRole < needRole{
		reply.Message = "当前用户无权拉取新用户"
		return reply,nil
	}
	err = mapper.AddUserToChat(chatId, request.UserId)
	if err == nil{
		reply.Success = true
	} else{
		reply.Message = err.Error()
	}
	return reply,err
}
