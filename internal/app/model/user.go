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
	ID           int64  `xorm:"pk autoincr 'id'" json:"id"`
	Username     string `xorm:"'username'" json:"username"`
	PasswordHash string `xorm:"'password_hash'" json:"-"`
	RoleID       int64  `xorm:"'role_id'" json:"role_id"`
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
