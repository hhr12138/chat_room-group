package object

type BasicItem struct {
	UserId int64 `from:"user_id",binding:"required"`
}
