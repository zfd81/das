package dao

import "github.com/zfd81/rooster/types/container"

type ConnectionDao interface {
	Save(entity *ConnectionInfo) error
	DeleteById(id string) error
	Modify(entity *ConnectionInfo) error
	FindById(id string) (container.Map, error)
	FindAllByProject(projectCode string) ([]container.Map, error)
}

type ConnectionDaoImpl struct {
}

func (c *ConnectionDaoImpl) Save(entity *ConnectionInfo) (err error) {
	_, err = db.Save(entity)
	return
}

func (c *ConnectionDaoImpl) DeleteById(id string) (err error) {
	sql := "delete from das_connection where conn_id=:val"
	_, err = db.Exec(sql, id)
	return
}

func (c *ConnectionDaoImpl) Modify(entity *ConnectionInfo) (err error) {
	sql := `
		UPDATE das_connection SET 
  			conn_name =: conn_name,
			driver =: driver,
			address =: address,
			port =: port,
			user_name =: user_name,
			password =: password,
			db =: db,
		 	modifier =: modifier,
			modified_time =: modified_time
		WHERE
			conn_id =: conn_id
	`
	_, err = db.Exec(sql, entity)
	return
}

func (c *ConnectionDaoImpl) FindById(id string) (container.Map, error) {
	sql := "select conn_id,conn_name,driver,address,port,user_name,password,db from das_connection where conn_id=:val"
	return db.QueryMap(sql, id)
}

func (c *ConnectionDaoImpl) FindAllByProject(projectCode string) ([]container.Map, error) {
	sql := "select conn_id,conn_name,driver,address,port,user_name,password,db,DATE_FORMAT(created_time, '%Y-%m-%d %H\\:%i') AS created_time,modified_time from das_connection where project_code=:val ORDER BY created_time"
	r, err := db.Query(sql, projectCode)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func NewConnectionDao() *ConnectionDaoImpl {
	return &ConnectionDaoImpl{}
}
