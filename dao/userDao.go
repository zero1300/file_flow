package dao

import (
	"context"
	"errors"
	"file_flow/dao/calc"
	"file_flow/ent"
	"file_flow/ent/mixin"
	"file_flow/ent/user"
	"file_flow/global"
	"file_flow/models"
	"fmt"
	"time"
)

var userDao *UserDao

type UserDao struct {
	db *ent.Client
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{db: global.Client}
	}
	return userDao
}

func (u UserDao) GetUserByEmail(email string) (*ent.User, error) {
	return u.db.User.Query().Where(user.Email(email)).First(context.Background())
}

func (u UserDao) CountUserByEmailSkipSoftDel(email string) (int, error) {
	i, err := u.db.User.Query().Where(user.Email(email)).Aggregate(ent.Count()).Int(mixin.SkipSoftDelete(context.Background()))
	return i, err
}

func (u UserDao) GetUserDeleted(email string) (*ent.User, error) {
	return u.db.User.Query().Where(user.Email(email), user.CreateAtLT(time.Now())).First(mixin.SkipSoftDelete(context.Background()))
}

func (u UserDao) CountEmail(email string) int {
	count := u.db.User.Query().Where(user.Email(email)).Aggregate(ent.Count()).IntX(context.Background())
	return count
}

func (u UserDao) AddUser(email, password, nickname string) error {

	userPo, err := u.GetUserDeleted(email)
	if err != nil {
		fmt.Println(err.Error())
	}
	if userPo.ID == 0 {
		_, err := u.db.User.Create().
			SetEmail(email).
			SetNickname(nickname).
			SetPassword(password).SetAvatar(" ").Save(context.Background())
		return err
	}
	_, err = u.db.User.UpdateOneID(userPo.ID).ClearDeleteAt().Save(context.Background())
	return err
}

func (u UserDao) GetUserById(id int) (*ent.User, error) {
	userPo, err := u.db.User.Get(context.Background(), id)
	if ent.IsNotFound(err) {
		err = errors.New("用户不存在")
	}
	return userPo, err
}

func (u UserDao) UserList(p models.Paginate) ([]*ent.User, int, error) {
	calc.Offset(&p)
	count := u.db.User.Query().CountX(context.Background())

	list, err := u.db.User.Query().Limit(p.PageSize).Offset(p.Page).Order(ent.Desc(user.FieldCreateAt)).All(context.Background())
	return list, count, err
}

func (u UserDao) UpdateUser(user models.UserUpdateModel) (*ent.User, error) {
	return u.db.User.UpdateOneID(user.ID).SetNickname(user.Nickname).SetAvatar(user.Avatar).Save(context.Background())
}

func (u UserDao) DelUser(id int) error {
	err := u.db.User.DeleteOneID(id).Exec(context.Background())
	return err
}
