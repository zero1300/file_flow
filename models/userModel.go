package models

type UserLoginModel struct {
	Email    string `json:"email" form:"email" binding:"required,len=11"`
	Password string `json:"-" form:"password" binding:"required"`
}

type UserRegModel struct {
	UserLoginModel
	Nickname string `form:"nickname" `
}

type UserUpdateModel struct {
	ID       int    `form:"id" binding:"required"`
	Nickname string `form:"nickname"`
	Avatar   string `form:"avatar"`
}
