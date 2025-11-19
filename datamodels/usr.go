package datamodels

type User struct {
	Id   int64 `sql:"userId" form:"userId" json:"userId"`
	NickName string `sql:"nickName" form:"nickName" json:"nickName"`
	UserName string `sql:"userName" form:"userName" json:"userName"`
	HashPassword string `json:"-" form:"password" sql:"password"`
}
