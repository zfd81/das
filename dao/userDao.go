package dao

import "github.com/zfd81/rooster/types/container"

type UserDao interface {
	FindByName(name string) (*UserInfo, error)
	FindByNameAndPwd(name string, pwd string) (container.Map, error)
	SaveUserProject(relations []container.JsonMap) error
	DeleteUserProject(relations []container.JsonMap) error
	FindAllByProject(projectCode string) ([]container.Map, error)
	FindAllNotInProject(uid string, projectCode string) ([]container.Map, error)
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) FindByName(name string) (*UserInfo, error) {
	user := &UserInfo{}
	sql := "select id from das_sys_user where name=:val"
	err := db.QueryStruct(user, sql, name)
	return user, err
}

func (u *UserDaoImpl) FindByNameAndPwd(name string, pwd string) (container.Map, error) {
	mp := map[string]interface{}{
		"name": name,
		"pwd":  pwd,
	}
	sql := "select id,name,full_name from das_sys_user where status='1' and name=:name and password=:pwd"
	return db.QueryMap(sql, mp)
}

func (u *UserDaoImpl) SaveUserProject(relations []container.JsonMap) (err error) {
	sql := `
		INSERT INTO das_rel_user_project (
			user_id,
			project_code,
			status,
			creator,
			created_time,
			modifier,
			modified_time
		)
		VALUES
			{@vals [,] (
				:this.id,
				:this.code ,
				'1',
				:this.uid ,
				:this.t,
				:this.uid ,
				:this.t
			) }
	`
	_, err = db.Exec(sql, relations)
	return
}

func (u *UserDaoImpl) DeleteUserProject(relations []container.JsonMap) (err error) {
	sql := `
		DELETE
		FROM
			das_rel_user_project
		WHERE
			{@vals [OR] (user_id=:this.id AND project_code=:this.code) }
	`
	_, err = db.Exec(sql, relations)
	return
}

func (c *UserDaoImpl) FindAllByProject(projectCode string) ([]container.Map, error) {
	sql := `
		SELECT DISTINCT
			user_id,
			NAME,
			full_name,
			phone_number,
			email,
			created_time
		FROM
			(
				SELECT
					user_id,
					name,
					full_name,
					phone_number,
					email,
					DATE_FORMAT(t2.created_time, '%Y-%m-%d %H\\:%i') AS created_time
				FROM
					das_sys_user t1,
					das_rel_user_project t2
				WHERE
					t1.id = t2.user_id
				AND t2.project_code = :val
			) t
	`
	r, err := db.Query(sql, projectCode)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func (c *UserDaoImpl) FindAllNotInProject(uid string, projectCode string) ([]container.Map, error) {
	mp := map[string]interface{}{
		"uid":         uid,
		"projectCode": projectCode,
	}
	sql := `
		SELECT
			id user_id,
			NAME,
			full_name,
			phone_number,
			email
		FROM
			das_sys_user
		WHERE
			id NOT IN (
				SELECT
					user_id
				FROM
					das_rel_user_project
				WHERE
					project_code = :projectCode
			)
		AND id <> :uid
	`
	r, err := db.Query(sql, mp)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{}
}
