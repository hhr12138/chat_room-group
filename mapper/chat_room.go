package mapper

import (
	"database/sql"
	"github.com/hhr12138/chat_room-group/object"
	"github.com/jmoiron/sqlx"
	"time"
)

type ChatRoom struct {
	Id          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Introduce   sql.NullString `db:"introduce"`
	HeadImage   sql.NullString `db:"head_image"`
	AddUserRole sql.NullInt64  `db:"add_user_role"`
	Del         sql.NullBool   `db:"del"`
	GmtCreate   sql.NullInt64  `db:"gmt_create"`
	GmtModified sql.NullInt64  `db:"gmt_modified"`
}

const (
	CHAT_ROOM = "`chat_room`"
)

func GetAddUserRoleById(id int64) (int64, error) {
	var result int64
	sql := "select `add_user_role` from " + CHAT_ROOM + " where id=? and `del`=?"
	err := object.Db.Get(result, sql, id, false)
	return result, err
}

func DeleteChatRoom(id int64) error {
	sql := "update `del`=true from `chat_room` where `id`=id"
	_, err := object.Db.Exec(sql, id)
	return err
}

func InsertChatRoom(chatRoom *ChatRoom) (int64, error) {
	now := time.Now().UnixMilli()
	sql := "insert into `chat_room`(`name`,`introduce`,`del`,`gmt_create`,`gmt_modified`) values(?,?,?,?,?)"
	exec, err := object.Db.Exec(sql, chatRoom.Name, chatRoom.Introduce, false, now, now)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	return id, err
}

func SelectGroupsByUserId(groupIds []int64) ([]*ChatRoom, error) {
	resp := make([]*ChatRoom,0)
	q, args, err := sqlx.In("select `id`, `name`, `head_image` from "+CHAT_ROOM+" where `id` in (?) and !del", groupIds)
	err = object.Db.Select(&resp,q,args...)
	return resp, err
}
