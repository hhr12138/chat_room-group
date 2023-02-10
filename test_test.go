package chat_room

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door"
	"github.com/hhr12138/door/object"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestRegister(t *testing.T) {
	db, _ := sqlx.Open("mysql", object.TEST_DB_ADDR)
	engine := gin.Default()
	door.Register(db,engine)
	Register(engine,db)
	engine.Run("127.0.0.1:8080")
}
