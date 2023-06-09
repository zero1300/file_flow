package models

type UserLoginModel struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required"`
}

type UserRegModel struct {
	UserLoginModel
	Nickname string `form:"nickname" `
	Code     string `form:"code" binding:"required,len=4"`
}

type UserUpdateModel struct {
	ID       int    `form:"id" binding:"required"`
	Nickname string `form:"nickname"`
	Avatar   string `form:"avatar"`
}
