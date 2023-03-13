package api

import (
	"file_flow/common/resp"
	"file_flow/models"
	"file_flow/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
	userService *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{userService: service.NewUserService()}
}

func (u UserApi) Login(ctx *gin.Context) {
	var loginUserModel models.UserLoginModel
	err := ctx.ShouldBind(&loginUserModel)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	token, err := u.userService.Login(loginUserModel.Email, loginUserModel.Password)
	if err != nil {
		resp.Fail(ctx, "登录失败: "+err.Error())
		return
	}

	var data = gin.H{
		"token": token,
	}
	resp.Success(ctx, data)
}

func (u UserApi) Register(ctx *gin.Context) {
	var form models.UserRegModel

	err := ctx.ShouldBind(&form)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}

	if form.Nickname == "" {
		form.Nickname = form.Email
	}

	err = u.userService.AddUser(form.Email, form.Password, form.Nickname)
	if err != nil {
		resp.Fail(ctx, "注册失败: "+err.Error())
		return
	}
	resp.Success(ctx, gin.H{"msg": "注册成功"})
}

func (u UserApi) GetUserById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	userEnt, err := u.userService.GetUserById(id)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, userEnt)
}

func (u UserApi) UserList(ctx *gin.Context) {
	paginate, err := models.GetPaginate(ctx)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	list, count, err := u.userService.UserList(paginate)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, gin.H{
		"users": list,
		"total": count,
	})
}

func (u UserApi) UpdateUser(ctx *gin.Context) {
	var userUpdateModel models.UserUpdateModel
	err := ctx.ShouldBind(&userUpdateModel)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	user, err := u.userService.UpdateUser(userUpdateModel)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, user)
}

func (u UserApi) DelUser(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	err := u.userService.DelUser(id)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	resp.Success(ctx, nil)
}
