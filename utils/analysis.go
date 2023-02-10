package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hhr12138/door/entity"
)

func GetUser(ctx *gin.Context) *entity.User{
	u,ok := ctx.Get("user")
	if !ok{
		return new(entity.User)
	}
	user,ok := u.(*entity.User)
	if !ok{
		return new(entity.User)
	}
	return user
}
