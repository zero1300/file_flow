package service

import (
	"context"
	"errors"
	"file_flow/common/jwt"
	"file_flow/dao"
	"file_flow/ent"
	"file_flow/global"
	"file_flow/global/constants"
	"file_flow/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var userService *UserService

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{userDao: dao.NewUserDao()}
	}
	return userService
}

func getKey(id int) string {
	return constants.TokenKey + strconv.Itoa(id)
}

func getTokenById(id int) (string, error) {
	key := getKey(id)
	cmd := global.Redis.Get(context.Background(), key)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	val := cmd.Val()
	return val, nil
}

func storeToken(id int, token string) error {
	expire := viper.GetDuration("jwt.tokenExpire") * time.Hour
	key := getKey(id)
	return global.Redis.Set(context.Background(), key, token, expire).Err()
}

func (u UserService) Login(email, password string) (string, error) {
	user, err := u.userDao.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("用户不存在")
	}
	if user.Password != password {
		return "", errors.New("密码错误")
	}
	token, err := getTokenById(user.ID)
	if err == nil {
		return token, nil
	}
	token, err = jwt.GenerateToken(user.ID, user.Nickname)
	err = storeToken(user.ID, token)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return token, err
}

func (u UserService) AddUser(email, password, nickname string) error {
	count := u.userDao.CountEmail(email)
	if count == 1 {
		return errors.New("该手机号已被注册")
	}
	err := u.userDao.AddUser(email, password, nickname)
	if err != nil {
		logrus.Info(err.Error())
		return err
	}
	return nil
}

func (u UserService) GetUserById(id int) (*ent.User, error) {
	userPo, err := u.userDao.GetUserById(id)
	if err != nil {
		return &ent.User{}, err
	}
	return userPo, nil
}

func (u UserService) UserList(p models.Paginate) ([]*ent.User, int, error) {
	return u.userDao.UserList(p)
}

func (u UserService) UpdateUser(user models.UserUpdateModel) (*ent.User, error) {
	userPo, err := u.userDao.GetUserById(user.ID)
	if err != nil {
		return nil, err
	}
	if user.Nickname == "" {
		user.Nickname = userPo.Nickname
	}
	if user.Avatar == "" {
		user.Avatar = userPo.Avatar
	}
	return u.userDao.UpdateUser(user)
}

func (u UserService) DelUser(id int) error {
	return u.userDao.DelUser(id)
}
