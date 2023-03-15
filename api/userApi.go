package api

import (
	"context"
	"file_flow/common/email"
	"file_flow/common/helper"
	"file_flow/common/resp"
	"file_flow/global"
	"file_flow/global/constants"
	"file_flow/models"
	"file_flow/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
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

func (u UserApi) GetCode(ctx *gin.Context) {
	var form struct {
		Email string `form:"email" binding:"required,email"`
	}
	err := ctx.ShouldBind(&form)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}
	code := helper.GenerateVerificationCode()
	log.Println("验证码: " + code)
	err = email.SendVerificationCode(form.Email, code)
	if err != nil {
		resp.Fail(ctx, "验证码发送失败: "+err.Error())
		return
	}
	global.Redis.Set(context.Background(), constants.VerificationCodeKeyPrefix+form.Email, code, 3*time.Minute)
	resp.Success(ctx, gin.H{})
}

func (u UserApi) Register(ctx *gin.Context) {
	var form models.UserRegModel

	err := ctx.ShouldBind(&form)
	if err != nil {
		resp.Fail(ctx, err.Error())
		return
	}

	val := global.Redis.Get(context.Background(), constants.VerificationCodeKeyPrefix+form.Email).Val()
	if val != form.Code {
		resp.Fail(ctx, "验证码错误")
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

func (u UserApi) GetUsrByToken(ctx *gin.Context) {
	id, err := helper.GetUserIdByToken(ctx)
	if err != nil {
		resp.Fail(ctx, "token 无效")
		return
	}
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
