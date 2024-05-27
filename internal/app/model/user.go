package model

import (
	"xorm.io/xorm"
)

type RoleType uint

const (
	RoleAdmin RoleType = iota
	RoleUser
)

const hasUser = false

type User struct {
	ID           int64  `xorm:"pk autoincr 'id'" `
	Username     string `xorm:"'username'"`
	PasswordHash string `xorm:"'password_hash'"`
	Lang         string `xorm:"'lang'"`
	RoleID       int64  `xorm:"'role_id'"`
}

func (u *User) TableName() string {
	return "user"
}

func CreateUser(engine *xorm.Engine, user *User) error {
	_, err := engine.Insert(user)
	return err
}

func GetUserByUsername(engine *xorm.Engine, username string) (*User, error) {
	user := &User{Username: username}
	has, err := engine.Get(user)
	if !has {
		return nil, nil
	}
	return user, err
}

func IfUsernameExist(engine *xorm.Engine, username string) (bool, error) {
	return engine.Exist(&User{Username: username})
}

func HasUser(engine *xorm.Engine) (bool, error) {
	return hasUser, nil
}

func UpdateUser(engine *xorm.Engine, user *User) error {
	_, err := engine.ID(user.ID).Update(user)
	return err
}

func GetUserByID(engine *xorm.Engine, id int64) (*User, error) {
	user := &User{ID: id}
	has, err := engine.Get(user)
	if !has {
		return nil, nil
	}
	return user, err
}
