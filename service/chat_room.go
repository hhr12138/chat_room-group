package service

import (
	"database/sql"
	"github.com/hhr12138/chat_room-group/mapper"
	"github.com/hhr12138/chat_room-group/object"
	"github.com/hhr12138/chat_room-group/vo"
	"github.com/hhr12138/chat_room-utils/change"
)

const (
	OWNER   = 0
	MANAGER = 1000
	USER    = 10000
)

type CreateChatRoomRequest struct {
	object.BasicItem
	Name      string `from:"name",binding:"required"`
	Introduce string `from:"introduce"`
}

type CreateChatRoomResponse struct {
	Id int64 `json:"id"`
}

type DeleteChatRoomRequest struct {
	object.BasicItem
	Id     int64 `from:"id",binding:"required"`
	UserId int64 `from:"id",binding:"required"`
}

type DeleteChatRoomResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func DeleteChatRoom(request *DeleteChatRoomRequest) (*DeleteChatRoomResponse, error) {
	reply := &DeleteChatRoomResponse{}
	ok, err := checkPermission(request, OWNER)
	if err != nil {
		return nil, err
	}
	if !ok {
		reply.Message = "只有群主可以删除群聊"
		return reply, nil
	}
	err = mapper.DeleteChatRoom(request.Id)
	if err != nil {
		return nil, err
	}
	reply.Success = true
	return reply, nil
}

func checkPermission(request *DeleteChatRoomRequest, target int64) (bool, error) {
	role, err := mapper.GetUserRoleInGroup(request.Id, request.UserId)
	if err != nil {
		return false, err
	}
	return role <= target, err
}

func CreateChatRoom(request *CreateChatRoomRequest) (*CreateChatRoomResponse, error) {
	chatRoom := &mapper.ChatRoom{
		Name: sql.NullString{
			String: request.Name,
			Valid:  true,
		},
		Introduce: sql.NullString{
			String: request.Introduce,
			Valid:  true,
		},
	}
	//建群和指定群主都执行成功后才认为成功
	room, err := mapper.InsertChatRoom(chatRoom)
	if err != nil {
		return nil, err
	}
	err = mapper.InsertUserToGroup(room, request.UserId, OWNER)
	if err != nil {
		return nil, err
	}
	reply := &CreateChatRoomResponse{
		Id: room,
	}
	return reply, nil
}

func GetAllUserId(chatId int64) ([]int64, error) {
	userIds, err := mapper.SelectUserIdsByChatId(chatId)
	return userIds, err
}

func GetGroupIdsByUserId(uid int64) ([]int64, error) {
	groupIds, err := mapper.SelectGroupIdsByUserId(uid)
	return groupIds, err
}

func GetGroupsByGroupId(groupIds []int64) ([]*vo.ChatRoom, error) {
	groups, err := mapper.SelectGroupsByUserId(groupIds)
	resp := make([]*vo.ChatRoom, 0)
	for _, group := range groups {
		chatRoom := new(vo.ChatRoom)
		change.SqlObjToObj(group, chatRoom)
		resp = append(resp, chatRoom)
	}
	return resp, err
}
