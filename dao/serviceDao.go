package dao

import "github.com/zfd81/rooster/types/container"

type ServiceDao interface {
	Save(entity *ServiceInfo) error
	DeleteById(id string) error
	Modify(entity *ServiceInfo) error
	ModifyName(entity *ServiceInfo) error
	ModifyCatalog(entity *ServiceInfo) error
	ModifyStatus(entity *ServiceInfo) error
	FindById(id, version string) (container.Map, error)
	FindAllById(id string) ([]container.Map, error)
	FindAllByCatalog(catalog string) ([]container.Map, error)
}

type ServiceDaoImpl struct {
}

func (s *ServiceDaoImpl) Save(entity *ServiceInfo) (err error) {
	_, err = db.Save(entity)
	return
}

func (s *ServiceDaoImpl) DeleteById(id string) (err error) {
	sql := "delete from das_service where serv_code=:val"
	_, err = db.Exec(sql, id)
	return
}

func (s *ServiceDaoImpl) Modify(entity *ServiceInfo) (err error) {
	sql := `
		UPDATE das_service SET
			serv_sql =:serv_sql,
			serv_param =:serv_param,
		 	modifier =:modifier,
			modified_time =:modified_time
		WHERE
			serv_code =:serv_code
		AND version =:version
	`
	_, err = db.Exec(sql, entity)
	return
}

func (s *ServiceDaoImpl) ModifyName(entity *ServiceInfo) (err error) {
	sql := `
		UPDATE das_service SET 
  			serv_name =:serv_name,
		 	modifier =:modifier,
			modified_time =:modified_time
		WHERE
			serv_code =:serv_code
		AND version =:version
	`
	_, err = db.Exec(sql, entity)
	return
}

func (s *ServiceDaoImpl) ModifyCatalog(entity *ServiceInfo) (err error) {
	sql := `
		UPDATE das_service SET 
  			serv_catalog =:serv_catalog,
		 	modifier =:modifier,
			modified_time =:modified_time
		WHERE
			serv_code =:serv_code
		AND version =:version
	`
	_, err = db.Exec(sql, entity)
	return
}

func (s *ServiceDaoImpl) ModifyStatus(entity *ServiceInfo) (err error) {
	sql := `
		UPDATE das_service SET 
  			status =:status,
		 	modifier =:modifier,
			modified_time =:modified_time
		WHERE
			serv_code =:serv_code
		AND version =:version
	`
	_, err = db.Exec(sql, entity)
	return
}

func (s *ServiceDaoImpl) FindById(id, version string) (container.Map, error) {
	sql := `
		SELECT
			serv_code,
			serv_name,
			serv_catalog,
			serv_type,
			serv_sql,
			serv_param,
			version
		FROM
			das_service
		WHERE
			serv_code =: val
		`
	return db.QueryMap(sql, id)
}

func (s *ServiceDaoImpl) FindAllById(id string) ([]container.Map, error) {
	sql := `
		SELECT
			serv_code,
			serv_name,
			serv_catalog,
			serv_type,
			serv_sql,
			serv_param,
			version,
			STATUS,
			creator,
			created_time,
			modifier,
			modified_time
		FROM
			das_service
		WHERE
			serv_code =:val
	`
	r, err := db.Query(sql, id)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func (s *ServiceDaoImpl) FindAllByCatalog(catalog string) ([]container.Map, error) {
	sql := `
		SELECT
			serv_code,
			MAX(serv_name),
			MAX(serv_type),
			MAX(version)
		FROM
			das_service
		WHERE
			STATUS = '1'
		AND serv_catalog =:val
		GROUP BY
			serv_code
	`
	r, err := db.Query(sql, catalog)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func NewServiceDao() *ServiceDaoImpl {
	return &ServiceDaoImpl{}
}
