package mapper

import (
	"github.com/hhr12138/chat_room/object"
	"time"
)

const USER_CHAT_ROOM = "`tb_user_chat_room`"

func AddUserToChat(chatId, userId int64) error {
	now := time.Now().UnixMilli()
	sql := "insert into " + USER_CHAT_ROOM + "(`chat_id`,`user_id`,`offset`,`del`,`gmt_create`,`gmt_modified`) " +
		"select (?,?,?,?,?,?) from dual where not exist " +
		"(select 1 from " + USER_CHAT_ROOM + " where `chat_id`=? and `user_id`=?)"
	_, err := object.Db.Exec(sql, chatId, userId, -1, false, now, now, chatId, userId)
	return err
}

func SelectUserIdsByChatId(chatId int64) ([]int64, error) {
	sql := "select `user_id` from " + USER_CHAT_ROOM + " where `chat_id`=? and `del`=false group by `chat_id`"
	reply := make([]int64, 0)
	err := object.Db.Select(reply, sql, chatId)
	return reply, err
}

func SelectGroupIdsByUserId(uid int64)([]int64, error){
	sql := "select `chat_id` from " + USER_CHAT_ROOM + " where `user_id`=? and `del`=false"
	reply := make([]int64,0)
	err := object.Db.Select(&reply,sql,uid)
	return reply,err
}

func GetUserRoleInGroup(chatId, userId int64) (int64, error) {
	sql := "select `role` from " + USER_CHAT_ROOM + " where `chat_id`=? and `user_id`=? and `del`=false"
	var reply int64
	err := object.Db.Get(reply, sql, chatId, userId)
	return reply, err
}

func InsertUserToGroup(chatId, userId int64, role int64) error {
	now := time.Now().UnixMilli()
	sql := "insert into " + USER_CHAT_ROOM + "(`chat_id`,`user_id`,`role`,`del`,`gmt_create`,`gmt_modified`) values(?,?,?,?,?,?)"
	_, err := object.Db.Exec(sql, chatId, userId, role, false, now, now)
	return err
}
