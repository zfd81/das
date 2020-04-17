package dao

import "github.com/zfd81/rooster/types/container"

type UserDao interface {
	FindByName(name string) (*UserInfo, error)
	FindByNameAndPwd(name string, pwd string) (container.Map, error)
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) FindByName(name string) (*UserInfo, error) {
	user := &UserInfo{}
	sql := "select id from sys_user where name=:val"
	err := db.QueryStruct(user, sql, name)
	return user, err
}

func (u *UserDaoImpl) FindByNameAndPwd(name string, pwd string) (container.Map, error) {
	mp := map[string]interface{}{
		"name": name,
		"pwd":  pwd,
	}
	sql := "select id,name,full_name from sys_user where status='1' and name=:name and password=:pwd "

	return db.QueryMap(sql, mp)
}

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{}
}
